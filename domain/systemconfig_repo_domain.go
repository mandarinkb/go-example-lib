package domain

import "github.com/mandarinkb/go-example-lib/model"

type SystemConfigRepo interface {
	GetSystemConfig(*[]model.SystemConfigValueModel) error
}
