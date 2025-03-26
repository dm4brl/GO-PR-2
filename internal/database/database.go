package database

import (
	"log"

	"github.com/dm4brl/GO-PR-2/internal/models"
	"gorm.io/driver/sqlite" // Или другой драйвер базы данных
	"gorm.io/gorm"
)

var DB *gorm.DB

// SetupDatabase - функция для настройки базы данных
func SetupDatabase() {
	var err error
	// Подключаемся к базе данных
	DB, err = gorm.Open(sqlite.Open("switch_status.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}

	// Выполняем миграцию
	err = DB.AutoMigrate(&models.SwitchStatus{})
	if err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	log.Println("База данных настроена и миграция выполнена")
}

// GetDB - вспомогательная функция для получения экземпляра базы данных
func GetDB() *gorm.DB {
	return DB
}
