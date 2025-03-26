package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/dm4brl/GO-PR-2/internal/config"
	"github.com/dm4brl/GO-PR-2/internal/database"
	"github.com/dm4brl/GO-PR-2/internal/models"
	"github.com/dm4brl/GO-PR-2/internal/scheduler"
	"github.com/dm4brl/GO-PR-2/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Загружаем конфигурацию
	config.LoadConfig()

	// Настройка базы данных
	database.SetupDatabase()

	// Настройка MQTT
	services.SetupMQTT(config.AppConfig.MQTTBroker)

	// Настройка Redis
	services.SetupRedis(config.AppConfig.RedisURL)

	// Запуск планировщика
	scheduler.StartScheduler()

	// Инициализация Gin
	router := gin.Default()

	// Пример маршрута
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Сервер запущен!"})
	})
	router.POST("/api/switch/status", func(c *gin.Context) {
		var status models.SwitchStatus
		// Привязываем JSON-пейлоад к структуре SwitchStatus
		if err := c.ShouldBindJSON(&status); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Здесь можно добавить сохранение в базу или логирование
		log.Printf("Получен статус свитча: ID=%s, State=%v, Timestamp=%d",
			status.ID, status.State, status.Timestamp)

		// Если требуется, можно сохранить в Redis/Postgres или передать в планировщик
		c.JSON(http.StatusOK, gin.H{"status": "received"})
	})
	// Запуск сервера на порту, указанном в конфигурации
	port := config.AppConfig.ServerPort
	log.Printf("Сервер запущен на порту %s", port)
	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
