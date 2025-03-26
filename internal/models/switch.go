package models

import "time"

// SwitchStatus представляет статус свитча, отправляемый эмулятором
type SwitchStatus struct {
	ID        string    `json:"id" gorm:"primaryKey" binding:"required"`
	State     bool      `json:"state" binding:"required"`
	Timestamp int64     `json:"timestamp" binding:"required"`
	UpdatedAt time.Time `json:"updated_at"`
}
