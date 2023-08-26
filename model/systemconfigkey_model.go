package model

type StaticSystemConfigKeyModel struct {
	API_CORE_PHP_ENDPOINT string
	S3_ACCESS_KEY_ID      string
	S3_SECRET_ACCESS_KEY  string
	S3_BUCKET_NAME        string
	S3_BUCKET             string
	S3_REGION             string
	S3_CDN_ENDPOINT       string

	SYSTEM_AUTH_SETTING_KEY string
	API_AUTHEN_CLIENT_KEY   string
	API_AUTHEN_ENDPOINT     string
	URL_ROOT_BOF            string
	BOF_ENDPOINT            string

	GUEST_ACCESS_TOKEN_PRIVATE_KEY                          string
	GUEST_ACCESS_TOKEN_PUBLIC_KEY                           string
	MEMBER_ACCESS_TOKEN_PRIVATE_KEY                         string
	MEMBER_ACCESS_TOKEN_PUBLIC_KEY                          string
	MEMBER_REFRESH_TOKEN_PRIVATE_KEY                        string
	MEMBER_REFRESH_TOKEN_PUBLIC_KEY                         string
	FIREBASE_SERVICE_ACCOUNT_PROJECT_NUMBER                 string
	FIREBASE_SERVICE_ACCOUNT_PROJECT_ID                     string
	FIREBASE_SERVICE_ACCOUNT_WEB_MOCK_FOR_ADMIN_USER_ID     string
	FIREBASE_SERVICE_ACCOUNT_WEB_MOCK_FOR_ADMIN_WEB_API_KEY string
	FIREBASE_SERVICE_ACCOUNT_WEB_MOCK_FOR_ADMIN_WEB         string
	FIREBASE_SERVICE_ACCOUNT_IOS                            string
	FIREBASE_SERVICE_ACCOUNT_ANDROID                        string

	DB_READ_WRITE_HOST         string
	DB_READ_WRITE_PORT         string
	DB_READ_WRITE_NAME         string
	DB_READ_WRITE_USERNAME_BOF string
	DB_READ_WRITE_PASSWORD_BOF string
	DB_READ_WRITE_USERNAME_JOB string
	DB_READ_WRITE_PASSWORD_JOB string
	DB_READ_ONLY_HOST          string
	DB_READ_ONLY_PORT          string
	DB_READ_ONLY_NAME          string
	DB_READ_ONLY_USERNAME_BOF  string
	DB_READ_ONLY_PASSWORD_BOF  string
	DB_READ_ONLY_USERNAME_JOB  string
	DB_READ_ONLY_PASSWORD_JOB  string

	REDIS_HOST     string
	REDIS_PORT     string
	REDIS_PASSWORD string
	REDIS_DB_NUM   string

	ENDPOINT_AUTHEN_HYPCODE string
	ENDPOINT_HYPCODE        string
	ENDPOINT_LIFESTYLE      string

	S3_CEGID_ACCESS_KEY_ID     string
	S3_CEGID_SECRET_ACCESS_KEY string
	S3_CEGID_REGION            string
	S3_CEGID_BUCKET            string
}
