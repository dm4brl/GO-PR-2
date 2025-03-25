package models

import "gorm.io/gorm"

type Schedule struct {
	gorm.Model
	DeviceID  uint   // Связь с устройством
	StartTime string `gorm:"not null"`
	EndTime   string `gorm:"not null"`
	DayOfWeek string `gorm:"not null"` // Дни недели для расписания
}
