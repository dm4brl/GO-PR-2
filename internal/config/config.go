package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort  string `mapstructure:"server_port"`
	DatabaseURL string `mapstructure:"database_url"`
	RedisURL    string `mapstructure:"redis_url"`
	MQTTBroker  string `mapstructure:"mqtt_broker"`
}

var AppConfig *Config

// LoadConfig загружает параметры конфигурации из файла с использованием viper.
func LoadConfig() {
	viper.SetConfigName("config") // Название конфигурационного файла
	viper.SetConfigType("yaml")   // Формат конфигурации
	viper.AddConfigPath(".")      // Путь для поиска файла конфигурации
	viper.AddConfigPath("./config")

	// Чтение конфигурации
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Ошибка чтения конфигурации: %v", err)
	}

	// Преобразуем конфигурацию в структуру
	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	log.Println("Конфигурация загружена успешно")
}

// GetDatabaseURL возвращает URL базы данных из конфигурации
func GetDatabaseURL() string {
	return AppConfig.DatabaseURL
}

// GetServerPort возвращает порт сервера из конфигурации
func GetServerPort() string {
	return AppConfig.ServerPort
}

// GetRedisURL возвращает URL для подключения к Redis
func GetRedisURL() string {
	return AppConfig.RedisURL
}

// GetMQTTBroker возвращает MQTT брокер из конфигурации
func GetMQTTBroker() string {
	return AppConfig.MQTTBroker
}
