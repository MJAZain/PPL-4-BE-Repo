package main

import (
	"go-gin-auth/config"
	"go-gin-auth/model"
	"go-gin-auth/router"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&model.User{})
	r := router.SetupRouter()
	r.Run(":8080")
}
