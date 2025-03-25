package services

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client
var Ctx = context.Background()

func SetupRedis(addr string) {
	Rdb = redis.NewClient(&redis.Options{
		Addr: addr,
	})

	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Ошибка подключения к Redis: %v", err)
	}
	log.Println("Redis подключен")
}
