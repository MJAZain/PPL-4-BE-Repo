package helpers

import (
	"go-gin-auth/config"
	"go-gin-auth/internal/category"
	"go-gin-auth/internal/incomingProducts"
	"go-gin-auth/internal/outgoingProducts"
	"go-gin-auth/internal/product"
	"go-gin-auth/internal/stock"
	"go-gin-auth/internal/unit"
	"go-gin-auth/model"
)

func MigrateDB() error {
	db := config.DB

	err := db.AutoMigrate(
		&model.User{},
		&model.ActivityLog{},
		&model.SystemConfig{},
		&product.Product{},
		&unit.Unit{},
		&category.Category{},
		&model.AuditLog{},
		&model.Transaksi{},
		&model.StockOpname{},
		&model.StockOpnameDetail{},
		&incomingProducts.IncomingProduct{},
		&incomingProducts.IncomingProductDetail{},
		&stock.Stock{},
		&outgoingProducts.OutgoingProduct{},
		&outgoingProducts.OutgoingProductDetail{},
	)

	if err != nil {
		return err
	}

	return nil
}
