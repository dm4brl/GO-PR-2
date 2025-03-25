package database

import (
	"log"

	"github.com/dm4brl/GO-PR-2/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabase() {
	var err error

	// Если хочешь использовать PostgreSQL, поменяй этот блок на:
	// dsn := config.AppConfig.DatabaseURL
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Device{}, &models.Schedule{})
	if err != nil {
		log.Fatalf("Ошибка миграции базы данных: %v", err)
	}

	log.Println("База данных настроена и миграция выполнена")
}
