package services

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var Client mqtt.Client

func SetupMQTT(broker string) {
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID("go-client")
	Client = mqtt.NewClient(opts)

	if token := Client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Ошибка подключения к MQTT: %v", token.Error())
	}
	log.Println("MQTT подключен")
}

func Publish(topic, message string) {
	if token := Client.Publish(topic, 0, false, message); token.Wait() && token.Error() != nil {
		log.Fatalf("Ошибка публикации MQTT сообщения: %v", token.Error())
	}
}
