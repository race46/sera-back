package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sera-back/models"
)

var Connection *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=sera port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		println("could not connected to db")
	}
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Basket{})
	Connection = db
}
