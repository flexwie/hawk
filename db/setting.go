package db

import (
	"hawk.wie.gg/models"
)

func GetVal(key string) string {
	var setting models.Setting
	DB.Model(&models.Setting{}).Where("key = ?", key).First(&setting)

	return setting.Value
}

func SetVal(key string, value string) {
	var setting models.Setting

	if DB.Model(&models.Setting{}).Where("key = ?", key).First(&setting); setting.Key == "" {
		DB.Save(&models.Setting{Key: key, Value: value})
	} else {
		DB.Model(&models.Setting{}).Where("key = ?", key).Updates(&models.Setting{Value: value})
	}
}
