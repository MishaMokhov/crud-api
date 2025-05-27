package database

import (
	"crud_api/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=!Krypto! dbname=crud_api port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed", err)
	}
	db.AutoMigrate(&models.User{}, &models.Product{},
		&models.ProductInWork{})
	DB = db
}
