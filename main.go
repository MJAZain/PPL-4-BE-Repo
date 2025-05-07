package main

import (
	"go-gin-auth/config"
	"go-gin-auth/model"
	"go-gin-auth/router"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&model.User{}, &model.ActivityLog{}, &model.SystemConfig{}, &model.AuditLog{}, &model.Transaksi{})
	r := router.SetupRouter()
	r.Run(":8080")
}
