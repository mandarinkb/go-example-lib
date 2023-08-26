package model

type StaticSystemConfigModel struct {
	ENV_APP_ENV                     string
	ENV_PROJECT_CODE                string
	ENV_SYSTEM_CONFIG_API_ENDPOINT  string
	ENV_SYSTEM_CONFIG_API_AUTHEN    string
	ENV_SYSTEM_CONFIG_API_PASSPHASE string

	ENV_DB_HOST     string
	ENV_DB_PORT     string
	ENV_DB_NAME     string
	ENV_DB_USERNAME string
	ENV_DB_PASSWORD string
}

type SystemConfigDBModel struct {
	DB_READ_WRITE_HOST     string
	DB_READ_WRITE_PORT     string
	DB_READ_WRITE_NAME     string
	DB_READ_WRITE_USERNAME string
	DB_READ_WRITE_PASSWORD string

	DB_READ_ONLY_HOST     string
	DB_READ_ONLY_PORT     string
	DB_READ_ONLY_NAME     string
	DB_READ_ONLY_USERNAME string
	DB_READ_ONLY_PASSWORD string
}

type SystemConfigResponseModel struct {
	Code string                `json:"code"`
	Data SystemConfigDataModel `json:"data"`
}

type SystemConfigDataModel struct {
	SystemConfig []SystemConfigDataDetailModel `json:"systemConfig"`
}

type SystemConfigDataDetailModel struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type GetSystemConfigResponseBodyModel struct {
	SystemConfig []GetSystemConfigResponseBodyDetailModel `json:"SystemConfig"`
}

type GetSystemConfigResponseBodyDetailModel struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
