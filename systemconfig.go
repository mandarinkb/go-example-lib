package systemconfig

import (
	"encoding/json"
	"fmt"
	"reflect"

	"io"
	"log"
	"net/http"
	"strings"

	"time"

	"github.com/mandarinkb/go-example-lib/model"
	"github.com/mandarinkb/go-example-lib/repository"
	"github.com/mandarinkb/go-example-lib/service"
	"github.com/mandarinkb/go-example-lib/util/config"
	"github.com/mandarinkb/go-example-lib/util/database"
	"github.com/mandarinkb/go-example-lib/util/encryption"
)

var (
	Key        model.StaticSystemConfigKeyModel
	staticKey  model.StaticSystemConfigModel
	expireOn   time.Time
	envOS      []string
	storeValue []model.GetSystemConfigResponseBodyDetailModel
)

// Get function require parameter key to get value and 1 optional to defined project otherwise it will be default
// 1. key [string] [require]
// 2. project code [string] [optional]

func Get(key string, args ...interface{}) string {
	if expireOn.Before(time.Now()) {
		updateSystemConfig()
	}

	// if slices.Index(envOS, key) != -1 {
	// 	return viper.GetString(key)
	// }
	// projectCode := viper.GetString(Key.ENV_PROJECT_CODE)
	projectCode := staticKey.ENV_PROJECT_CODE

	//args[0] == projectCode
	if len(args) > 0 {
		if pc, ok := args[0].(string); ok {
			projectCode = pc
		}
	}

	//TODO: set args
	value := ""
	key = fmt.Sprintf("%s_%s_%s", staticKey.ENV_APP_ENV, projectCode, key)
	for _, data := range storeValue {
		if strings.EqualFold(data.Key, key) {
			value = data.Value
			break
		}
	}
	return value
}

func Init(env model.StaticSystemConfigModel) error {
	// set systemconfigkey.Key เพื่อตอนเรียกใช้ Get(systemconfigkey.Key.BOF_ENDPOINT) ได้เลย
	fields := reflect.TypeOf(Key)
	for i := 0; i < fields.NumField(); i++ {
		key := fields.Field(i).Name
		reflect.ValueOf(&Key).Elem().FieldByName(key).SetString(key)
		envOS = append(envOS, key)
	}

	// check public / private
	switch {
	case strings.EqualFold(env.ENV_APP_ENV, "") || strings.EqualFold(env.ENV_PROJECT_CODE, "") || strings.EqualFold(env.ENV_SYSTEM_CONFIG_API_PASSPHASE, ""):
		return fmt.Errorf("not found app env or env project code or api passphase")
	case !strings.EqualFold(env.ENV_SYSTEM_CONFIG_API_ENDPOINT, "") && // case public
		!strings.EqualFold(env.ENV_SYSTEM_CONFIG_API_AUTHEN, "") &&
		!strings.EqualFold(env.ENV_SYSTEM_CONFIG_API_PASSPHASE, ""):
		return prepareValueFromPublic(env)
	case !strings.EqualFold(env.ENV_DB_HOST, "") && // case private
		!strings.EqualFold(env.ENV_DB_PORT, "") &&
		!strings.EqualFold(env.ENV_DB_NAME, "") &&
		!strings.EqualFold(env.ENV_DB_USERNAME, "") &&
		!strings.EqualFold(env.ENV_DB_PASSWORD, ""):
		return prepartValueFromPrivate(env)
	default:
		return fmt.Errorf("any get value for init error")
	}
}

func prepartValueFromPrivate(env model.StaticSystemConfigModel) error {
	staticKey.ENV_APP_ENV = env.ENV_APP_ENV
	staticKey.ENV_PROJECT_CODE = env.ENV_PROJECT_CODE
	staticKey.ENV_SYSTEM_CONFIG_API_PASSPHASE = env.ENV_SYSTEM_CONFIG_API_PASSPHASE

	staticKey.ENV_DB_HOST = env.ENV_DB_HOST
	config.Env.DB_READ_ONLY_HOST = env.ENV_DB_HOST
	config.Env.DB_READ_WRITE_HOST = env.ENV_DB_HOST

	staticKey.ENV_DB_PORT = env.ENV_DB_PORT
	config.Env.DB_READ_ONLY_PORT = env.ENV_DB_PORT
	config.Env.DB_READ_WRITE_PORT = env.ENV_DB_PORT

	staticKey.ENV_DB_NAME = env.ENV_DB_NAME
	config.Env.DB_READ_ONLY_NAME = env.ENV_DB_NAME
	config.Env.DB_READ_WRITE_NAME = env.ENV_DB_NAME

	staticKey.ENV_DB_USERNAME = env.ENV_DB_USERNAME
	config.Env.DB_READ_ONLY_USERNAME = env.ENV_DB_USERNAME
	config.Env.DB_READ_WRITE_USERNAME = env.ENV_DB_USERNAME

	staticKey.ENV_DB_PASSWORD = env.ENV_DB_PASSWORD
	config.Env.DB_READ_ONLY_PASSWORD = env.ENV_DB_PASSWORD
	config.Env.DB_READ_WRITE_PASSWORD = env.ENV_DB_PASSWORD

	database.InitDatabase()

	systemConf := model.GetSystemConfigResponseBodyModel{}
	// initil repo
	newSystemConfigRepo := repository.NewSystemConfigRepo()
	newSystemConfigService := service.NewSystemConfigService(newSystemConfigRepo)
	if err := newSystemConfigService.GetSystemConfig(&systemConf); err != nil {
		fmt.Println(err)
		return err
	}
	// fmt.Println(systemConf.SystemConfig)

	storeValue = nil // clear value
	sysDetail := model.GetSystemConfigResponseBodyDetailModel{}
	prefix := fmt.Sprintf("%s_%s", env.ENV_APP_ENV, env.ENV_PROJECT_CODE)
	for _, item := range systemConf.SystemConfig {
		//เอาไว้เช็ค กรณี env local/develop  มี response กลับมาทั้ง 2 environment
		if strings.Contains(item.Key, prefix) {
			sysDetail.Key = item.Key
			sysDetail.Value = decrypt(item.Value)
			storeValue = append(storeValue, sysDetail) // เก็บลง storeValue
		}
	}
	//expire 1 hour
	expireOn = time.Now().Add(time.Hour)
	return nil
}

func prepareValueFromPublic(env model.StaticSystemConfigModel) error {
	staticKey.ENV_APP_ENV = env.ENV_APP_ENV
	staticKey.ENV_PROJECT_CODE = env.ENV_PROJECT_CODE
	staticKey.ENV_SYSTEM_CONFIG_API_PASSPHASE = env.ENV_SYSTEM_CONFIG_API_PASSPHASE

	staticKey.ENV_SYSTEM_CONFIG_API_ENDPOINT = env.ENV_SYSTEM_CONFIG_API_ENDPOINT
	staticKey.ENV_SYSTEM_CONFIG_API_AUTHEN = env.ENV_SYSTEM_CONFIG_API_AUTHEN

	resp, err := callAPIGetSystemConfig()
	if err != nil {
		log.Println(err)
		return err
	}
	storeValue = nil // clear value
	sysDetail := model.GetSystemConfigResponseBodyDetailModel{}
	prefix := fmt.Sprintf("%s_%s", env.ENV_APP_ENV, env.ENV_PROJECT_CODE)
	for _, item := range resp.Data.SystemConfig {
		//เอาไว้เช็ค กรณี env local/develop  มี response กลับมาทั้ง 2 environment
		if strings.Contains(item.Key, prefix) {
			sysDetail.Key = item.Key
			sysDetail.Value = decrypt(item.Value)
			storeValue = append(storeValue, sysDetail) // เก็บลง storeValue
		}
	}

	//expire 1 hour
	expireOn = time.Now().Add(time.Hour)
	return nil
}

func callAPIGetSystemConfig() (model.SystemConfigResponseModel, error) {
	var response model.SystemConfigResponseModel

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, staticKey.ENV_SYSTEM_CONFIG_API_ENDPOINT+"/get", nil)
	if err != nil {
		log.Println(err)
		return response, err
	}
	apiAuthen := fmt.Sprintf("Bearer %s", staticKey.ENV_SYSTEM_CONFIG_API_AUTHEN)
	req.Header.Set("Authorization", apiAuthen)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Channel", "Job")

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response, err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return response, fmt.Errorf("system http status %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println(err)
		return response, err
	}

	if response.Code != "1" {
		log.Println(err)
		return response, fmt.Errorf("response code %d", resp.StatusCode)
	}
	return response, nil
}

func updateSystemConfig() {
	resp, err := callAPIGetSystemConfig()
	if err != nil {
		log.Println(err)
		return
	}
	storeValue = nil // clear value
	sysDetail := model.GetSystemConfigResponseBodyDetailModel{}
	prefix := fmt.Sprintf("%s_%s", staticKey.ENV_APP_ENV, staticKey.ENV_PROJECT_CODE)
	for _, item := range resp.Data.SystemConfig {
		//เอาไว้เช็ค กรณี env local/develop  มี response กลับมาทั้ง 2 environment
		if strings.Contains(item.Key, prefix) {
			sysDetail.Key = item.Key
			sysDetail.Value = decrypt(item.Value)
			storeValue = append(storeValue, sysDetail) // เก็บลง storeValue
		}
	}

	//expire 1 hour
	expireOn = time.Now().Add(time.Hour)
}

// getAll For test Only
func getAll() map[string]string {
	res := make(map[string]string)
	response, err := callAPIGetSystemConfig()
	if err != nil {
		log.Println(err)
		return res
	}

	// passphase := viper.GetString(Key.ENV_SYSTEM_CONFIG_API_PASSPHASE)
	passphase := staticKey.ENV_SYSTEM_CONFIG_API_PASSPHASE
	for i := 0; i < len(response.Data.SystemConfig); i++ {
		decrypted, err := encryption.CryptoJsAesDecrypt(passphase, response.Data.SystemConfig[i].Value)
		if err != nil {
			log.Println(err)
			return res
		}

		if valStr, ok := decrypted.(string); ok {
			res[response.Data.SystemConfig[i].Key] = valStr
		}
	}

	return res
}

// keyFromENV from field in ConfigModel
func decrypt(valueFromRespAPI string) string {
	pass := ""
	decrypted, err := encryption.CryptoJsAesDecrypt(staticKey.ENV_SYSTEM_CONFIG_API_PASSPHASE, valueFromRespAPI)
	if err != nil {
		fmt.Printf("Decrypt value error: %+v \n", err)
		return pass
	}

	if valStr, ok := decrypted.(string); ok {
		pass = valStr
	}
	return pass
}
