package main

import (
	"fmt"
	"log"

	"github.com/dm4brl/GO-PR-2/internal/config"
)

func main() {
	// Загружаем конфигурацию
	config.LoadConfig()

	// Запуск сервера
	fmt.Println("Сервер запущен на порту", config.AppConfig.ServerPort)

	// Здесь будет запуск HTTP-сервера
	log.Fatal("Сервер пока не реализован")
}
