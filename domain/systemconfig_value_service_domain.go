package domain

import "github.com/mandarinkb/go-example-lib/model"

type SystemConfigService interface {
	GetSystemConfig(*model.GetSystemConfigResponseBodyModel) error
}
