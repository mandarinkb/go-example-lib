package service

import (
	"github.com/mandarinkb/go-example-lib/domain"
	"github.com/mandarinkb/go-example-lib/model"
)

type systemConfigService struct {
	SystemConfigRepo domain.SystemConfigRepo
}

func NewSystemConfigService(systemConfigRepo domain.SystemConfigRepo) domain.SystemConfigService {
	return &systemConfigService{
		SystemConfigRepo: systemConfigRepo,
	}
}

func (service *systemConfigService) GetSystemConfig(systemConfig *model.GetSystemConfigResponseBodyModel) error {
	var systemConfigRepo []model.SystemConfigValueModel
	err := service.SystemConfigRepo.GetSystemConfig(&systemConfigRepo)
	if err != nil {
		// logg.LoggerGenerator.WriterLogging().Loggers.SetResponseTime(logg.LoggerGenerator.EndTime).Error("**********  Private :: GetSystemConfig APIs :: FAILED  **********", fmt.Sprintf("GetSystemConfig Failed :: %v %v", err.Error(), logg.GetCallerPathNameFileNameLineNumber()))
		return err
	}
	systemConfigDatas := []model.GetSystemConfigResponseBodyDetailModel{}

	//prepare data
	for i := 0; i < len(systemConfigRepo); i++ {
		systemConfigData := model.GetSystemConfigResponseBodyDetailModel{
			Key:   systemConfigRepo[i].Key,
			Value: systemConfigRepo[i].Value,
		}
		systemConfigDatas = append(systemConfigDatas, systemConfigData)
	}

	//response
	systemConfig.SystemConfig = systemConfigDatas

	return nil
}
