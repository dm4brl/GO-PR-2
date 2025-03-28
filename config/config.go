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

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// Добавляем несколько путей, где может быть конфигурация
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config") // если запуск из другой папки

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Ошибка чтения конфигурации: %v", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	log.Println("Конфигурация загружена успешно")
}
