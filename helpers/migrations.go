package helpers

import (
	"go-gin-auth/config"
	"go-gin-auth/internal/category"
	"go-gin-auth/internal/product"
	"go-gin-auth/internal/unit"
	"go-gin-auth/model"
)

func MigrateDB() error {
	db := config.DB

	err := db.AutoMigrate(
		&model.User{},
		&model.ActivityLog{},
		&model.AuditTrail{},
		&model.SystemConfig{},
		&product.Product{},
		&unit.Unit{},
		&category.Category{},
		&model.AuditLog{},
	)

	if err != nil {
		return err
	}

	return nil
}
