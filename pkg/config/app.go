package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	// Подключение к PostgreSQL
	dsn := "host=localhost user=postgres password=Hamid6997 dbname=books_store port=5432 sslmode=disable"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
