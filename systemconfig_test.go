package systemconfig

import (
	"fmt"

	"testing"

	"github.com/mandarinkb/go-example-lib/model"
)

func TestSystemconfig(t *testing.T) {
	//example case job or public
	env := model.StaticSystemConfigModel{
		ENV_APP_ENV:                     "LOCAL",
		ENV_PROJECT_CODE:                "LIFESTYLE",
		ENV_SYSTEM_CONFIG_API_PASSPHASE: "We@reTheBe$t^^",

		ENV_SYSTEM_CONFIG_API_ENDPOINT: "https://dev-lifestyle-api.deepblok.io/v1/icrm/systemconfig",
		ENV_SYSTEM_CONFIG_API_AUTHEN:   "uTvD3J8l5wD7LbkcxTqyuffYTkL0SZtsltZTsw481Og6mHaxJYdlSNw0RAQZj3AMVObNTNlN72UvAsgoiUnLNFOq32KzuksGPzC1GcVBXoLHCTaXcIOKW70GD2JxDUVZ06Iqun7LfGkY46fEpGtJ0lfYob6MCb4rcwK6w3iqPWIu0Y0DCWDvg8sZe89dwacT7aY4HNIjUqKwLI4kQOFKqMz22ZWqKEWMqHljK2C8oD0rLDIGGnXAFfH8VLROsN5ngEXZ4ZewRUApt48fCG842RNJJ1Umh0U7whqk4pbEdc4Eiu7gIvaXxNo3GyCBG6OA3bKuS5gbYwVo7f2HeiV1IZtPULaOgZRNdAphKWaY1kxTa4DSCvjMq0x3HHAE1fGzsRdE0Uf0rudwZdFB8AVw6EIHbl7DdPO7smyrxjbR7mZhDmONOxCf8oMhF0R6GkpdDmfDAwbDIPWgPsi8F0XxKp7gUCtafOLWcySuxNEJ6w8nmycdWhtm",
	}

	//example case private
	// env := model.StaticSystemConfigModel{
	// 	ENV_APP_ENV:                     "LOCAL",
	// 	ENV_PROJECT_CODE:                "LIFESTYLE",
	// 	ENV_SYSTEM_CONFIG_API_PASSPHASE: "We@reTheBe$t^^",

	// 	ENV_DB_HOST:     "none-prod-aurorav2.cluster-ro-c1usyj720z8q.ap-southeast-1.rds.amazonaws.com",
	// 	ENV_DB_PORT:     "3306",
	// 	ENV_DB_NAME:     "minorpluscrm_DEV",
	// 	ENV_DB_USERNAME: "usr.hongtonatthawee",
	// 	ENV_DB_PASSWORD: "MP9oj#TF8va$w6f",
	// }
	Init(env)

	fmt.Println(Get("ENDPOINT_HYPCODE"))
	fmt.Println(Get(Key.API_CORE_PHP_ENDPOINT))
	fmt.Println(Get(Key.GUEST_ACCESS_TOKEN_PRIVATE_KEY))
}
