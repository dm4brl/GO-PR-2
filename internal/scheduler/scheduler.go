package scheduler

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

// Dummy функция для проверки, является ли текущий день шаббатом или йом тов.
// Здесь можно подключить реальную логику на основе геолокации и еврейского календаря.
func CheckHoliday() bool {
	// Для примера: пусть шаббат начинается в субботу
	if time.Now().Weekday() == time.Saturday {
		return true
	}
	return false
}

func StartScheduler() {
	// Создаем новый cron с поддержкой секунд
	c := cron.New(cron.WithSeconds())

	// Добавляем задачу, которая будет выполняться каждый день в 00:00:00
	// Расписание можно настроить по необходимости.
	_, err := c.AddFunc("0 0 0 * * *", func() {
		log.Println("Ежедневная проверка запускается...")
		if CheckHoliday() {
			log.Println("Обнаружен шаббат или йом тов! Запускаем очередь задач...")
			// Здесь вызывайте функцию, которая формирует очередь задач для устройств
		} else {
			log.Println("Сегодня не шаббат/йом тов.")
		}
	})
	if err != nil {
		log.Fatalf("Ошибка планировщика: %v", err)
	}

	c.Start()
	log.Println("Планировщик запущен")
}
