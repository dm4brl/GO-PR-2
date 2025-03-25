package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string   `gorm:"unique;not null"`
	Email    string   `gorm:"unique;not null"`
	Location string   `gorm:"not null"`
	Devices  []Device `gorm:"foreignKey:UserID"` // Связь один ко многим
}
