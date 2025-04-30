// service/user_log_service.go
package service

import (
	"go-gin-auth/config"
	"go-gin-auth/model"
)

func CreateUserLog(userID uint, activity string) error {
	log := model.UserLog{
		UserID:   userID,
		Activity: activity,
	}
	return config.DB.Create(&log).Error
}

func GetAllUserLogs() ([]model.UserLog, error) {
	var logs []model.UserLog
	err := config.DB.Preload("User").Order("created_at desc").Find(&logs).Error
	return logs, err
}
