package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	math_rand "math/rand"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// var key = config.Env.ENCRYPTION_KEY
var key = "kN6ulzxbp29cKMGTw5lMyLMzdz7jkqn5"
var Rand *math_rand.Rand = math_rand.New(math_rand.NewSource(time.Now().UnixNano()))

const (
	charset      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	numberSet    = "0123456789"
	hexNumberSet = "0123456789abcdef"
)

func EncryptParamsValue(param string) string {
	k, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return ""
	}
	dataEncryption, _ := EncryptAES256GCMHex(param, k)
	if err != nil {
		return ""
	}
	return dataEncryption
}

func DecryptParamsValue(encryptString string) string {
	k, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return ""
	}
	dataDecryption, err := DecryptAES256GCMHex(encryptString, k)
	if err != nil {
		return ""
	}
	return dataDecryption
}

// EncryptAES256GCMHex is Encrypt AES256 GCM
func EncryptAES256GCMHex(data string, key []byte) (string, error) {
	text := []byte(data)

	if len(key) == 0 {
		errMsg := "Error Get Config Key"
		return "", fmt.Errorf("%v", errMsg)
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	seal := gcm.Seal(nil, nonce, text, nil)

	output := fmt.Sprintf("%x%x", nonce, seal)
	return output, nil
}

// DecryptAES256GCMHex is Encrypt AES256 GCM
func DecryptAES256GCMHex(encrypt string, key []byte) (string, error) {
	ciphertext, err := hex.DecodeString(encrypt)
	if err != nil {
		return "", err
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}
	nonceSize := gcm.NonceSize()

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func RandomString(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[Rand.Intn(len(charset))]
	}
	return string(result)
}

func RandomNumber(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = numberSet[Rand.Intn(len(numberSet))]
	}
	return string(result)
}

// EncryptBase64 is Encrypt base64
func EncryptBase64(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func RandomHexNumber(n int) string {
	bytes := make([]byte, n)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func HashAndSalt(data []byte) string {

	hash, err := bcrypt.GenerateFromPassword(data, bcrypt.MinCost)
	if err != nil {
		return ""
	}

	return string(hash)
}

func CompareHashAndSalt(hash string, data []byte) bool {

	byteHash := []byte(hash)
	if err := bcrypt.CompareHashAndPassword(byteHash, data); err != nil {
		return false
	}

	return true
}

func GetSHA256Hash(text string) string {
	hash := sha256.Sum256([]byte(text))
	return fmt.Sprintf("%x", hash)
}

func GenerateUniqueOrderID(key int, length int) string {
	keyBytes := []byte(fmt.Sprintf("%d", key))

	hash := sha256.New()
	hash.Write(keyBytes)
	hashBytes := hash.Sum(nil)

	uniqueString := base64.URLEncoding.EncodeToString(hashBytes)
	uniqueString = replaceNonAlphaNumeric(uniqueString)
	uniqueString = adjustStringLength(uniqueString, length)
	uniqueString = fmt.Sprintf("order_%d_%s", key, uniqueString)
	return uniqueString

}

func GenerateUniqueString(key string, length int) string {
	keyBytes := []byte(key)

	hash := sha256.New()
	hash.Write(keyBytes)
	hashBytes := hash.Sum(nil)

	uniqueString := base64.URLEncoding.EncodeToString(hashBytes)
	uniqueString = replaceNonAlphaNumeric(uniqueString)
	uniqueString = adjustStringLength(uniqueString, length)
	return uniqueString

}

func replaceNonAlphaNumeric(input string) string {
	replacer := strings.NewReplacer("+", "A", "/", "B", "=", "C", "-", "D", "_", "E")
	return replacer.Replace(input)
}

func adjustStringLength(input string, length int) string {
	if len(input) < length {
		input += strings.Repeat("0", length-len(input))
	} else if len(input) > length {
		input = input[:length]
	}

	return input
}

// CryptoJsAesEncrypt is func that encrypt data which has the same algorithm
// in Javascript function CryptoJS.aes.encrypt by using aes-256-cbc and PKCS5Padding
// passPhase is not key. it always generate new key from passphase
func CryptoJsAesEncrypt(passphrase string, value interface{}) (string, error) {
	salt := make([]byte, 8)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return "", err
	}

	salted := make([]byte, 0, 48)
	dx := make([]byte, 0)
	for len(salted) < 48 {
		hash := md5.New()
		hash.Write(dx)
		hash.Write([]byte(passphrase))
		hash.Write(salt)
		dx = hash.Sum(nil)
		salted = append(salted, dx...)
	}

	key := salted[:32]
	iv := salted[32:48]

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	encrypter := cipher.NewCBCEncrypter(block, iv)

	m, _ := json.Marshal(value)
	paddedValue := PKCS5Padding(m, aes.BlockSize)

	encryptedData := make([]byte, len(paddedValue))
	encrypter.CryptBlocks(encryptedData, paddedValue)

	result := make(map[string]string)
	result["ct"] = base64.StdEncoding.EncodeToString(encryptedData)
	result["iv"] = hex.EncodeToString(iv)
	result["s"] = hex.EncodeToString(salt)

	encrypted, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	b64 := base64.StdEncoding.EncodeToString(encrypted)
	return b64, nil
}

// PKCS5Padding for func CryptoJsAesEncrypt
func PKCS5Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// CryptoJsAesDecrypt is func that decrypt data which has the same algorithm
// in Javascript function CryptoJS.aes.decrypt by using aes-256-cbc and PKCS5Unpadding
// passPhase is not key. it always generate new key from passphase
func CryptoJsAesDecrypt(passphrase string, base64Encode string) (interface{}, error) {
	base64Decode, err := base64.StdEncoding.DecodeString(base64Encode)
	if err != nil {
		return "", err
	}

	var encrypted map[string]string
	err = json.Unmarshal(base64Decode, &encrypted)
	if err != nil {
		return "", err
	}

	s, err := hex.DecodeString(encrypted["s"])
	if err != nil {
		return "", err
	}

	ct, err := base64.StdEncoding.DecodeString(encrypted["ct"])
	if err != nil {
		return "", err
	}

	iv, err := hex.DecodeString(encrypted["iv"])
	if err != nil {
		return "", err
	}

	concatenatedPassphrase := passphrase + string(s)
	md5Data := make([][16]byte, 3)

	md5Data[0] = md5.Sum([]byte(concatenatedPassphrase))
	result := md5Data[0][:]
	for i := 1; i < 3; i++ {
		md5Data[i] = md5.Sum(append(md5Data[i-1][:], []byte(concatenatedPassphrase)...))
		result = append(result, md5Data[i][:]...)
	}

	key := result[:32]

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	object := make([]byte, len(ct))
	mode.CryptBlocks(object, ct)
	object = PKCS5Unpadding(object)

	var decrypted interface{}
	err = json.Unmarshal(object, &decrypted)
	if err != nil {
		return "", err
	}

	return decrypted, nil
}

// PKCS5Unpadding for func CryptoJsAesDecrypt
func PKCS5Unpadding(data []byte) []byte {
	padding := int(data[len(data)-1])
	return data[:len(data)-padding]
}

func DecodeBase64(value string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", fmt.Errorf("%v base64 could not decode key!, DecodeBase64 Function error!, %v", value, err.Error())

	}

	return string(decoded), nil
}

func MakeMD5(data string) string {

	// Generate the MD5 hash
	hasher := md5.New()
	hasher.Write([]byte(data))
	result := hex.EncodeToString(hasher.Sum(nil))

	return result
}

func EncodeBase64(value string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	return string(encoded)
}
