package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=ep-red-hat-a1aavee2-pooler.ap-southeast-1.aws.neon.tech user=neondb_owner password=npg_Z5pgeUSE2rwA dbname=neondb port=5432 sslmode=require"
	//dsn := "host=localhost user=postgres password=admin123 dbname=db port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	log.Println("Database connected")
}
