package database

import "gorm.io/gorm"

// SwitchStatus представляет модель для хранения статуса свитча в базе данных
type SwitchStatus struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	State     bool
	Timestamp int64
}
