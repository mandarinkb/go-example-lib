package repository

import (
	"github.com/mandarinkb/go-example-lib/domain"
	"github.com/mandarinkb/go-example-lib/model"
	"github.com/mandarinkb/go-example-lib/util/database"
)

type systemConfigRepo struct{}

func NewSystemConfigRepo() domain.SystemConfigRepo {
	return &systemConfigRepo{}
}

func (repo *systemConfigRepo) GetSystemConfig(systemConfig *[]model.SystemConfigValueModel) error {
	if err := database.Conn.Table(
		model.SystemConfig.TableName(),
	).Find(&systemConfig).Error; err != nil {
		return err
	}
	return nil
}
