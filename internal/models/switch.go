package models

import "time"

// SwitchStatus определяет данные свитча для базы данных.
type SwitchStatus struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	State     bool      `json:"state"`
	Timestamp int64     `json:"timestamp"`
	UpdatedAt time.Time `json:"updated_at"`
}
