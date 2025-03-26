package database

import (
	"log"

	"github.com/dm4brl/GO-PR-2/internal/config"
	"github.com/dm4brl/GO-PR-2/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// SetupDatabase устанавливает соединение с базой данных и выполняет миграции.
func SetupDatabase() {
	var err error
	// Используем PostgreSQL, строка подключения из конфигурации
	dsn := config.GetDatabaseURL() // Получаем строку подключения из конфигурации
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Выполняем миграцию для модели SwitchStatus
	if err := DB.AutoMigrate(&models.SwitchStatus{}); err != nil {
		log.Fatalf("Ошибка миграции базы данных: %v", err)
	}

	log.Println("База данных настроена и миграция выполнена")
}
