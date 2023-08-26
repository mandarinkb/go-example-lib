package model

import "time"

var SystemConfig SystemConfigValueModel

func (m *SystemConfigValueModel) TableName() string {
	return "System_ConfigValue"
}

type SystemConfigValueModel struct {
	ID              int64     `gorm:"column:ID;AUTO_INCREMENT"`
	SystemBrandCode string    `gorm:"column:SystemBrandCode"`
	Key             string    `gorm:"column:Key"`
	Value           string    `gorm:"column:Value"`
	DateCreate      time.Time `gorm:"column:DateCreate"`
	DateModify      time.Time `gorm:"column:DateModify"`
	Remark          string    `gorm:"column:Remark"`
}
