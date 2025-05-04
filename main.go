package main

import (
	"go-gin-auth/config"
	"go-gin-auth/model"
	"go-gin-auth/router"
	"os"

	product "go-gin-auth/internal/product"

	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		panic("error loading .env file: " + envErr.Error())
	}

	config.ConnectDB()

	migrateErr := config.DB.AutoMigrate(
		&model.User{},
		&model.ActivityLog{},
		&model.AuditTrail{},
		&model.SystemConfig{},
		&product.Product{},
	)
	if migrateErr != nil {
		panic("error migrating database: " + migrateErr.Error())
	}

	r := router.SetupRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	runErr := r.Run(":" + port)
	if runErr != nil {
		panic("error running server: " + runErr.Error())
	}
}
