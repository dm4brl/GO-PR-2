// internal/models/switch.go
package models

import "time"

// SwitchStatus определяет структуру для хранения данных о состоянии свитча.
type SwitchStatus struct {
	ID        string    `gorm:"primaryKey"`
	State     bool      `json:"state"`
	Timestamp time.Time `json:"timestamp"`
}
