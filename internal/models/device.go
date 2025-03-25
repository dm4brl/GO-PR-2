package models

import "gorm.io/gorm"

type Device struct {
	gorm.Model
	Name       string `gorm:"not null"`
	DeviceType string `gorm:"not null"`
	UserID     uint   // Идентификатор пользователя
	ScheduleID uint   // Связь с расписанием
}
