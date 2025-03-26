package models

// SwitchStatus представляет статус свитча, отправляемый эмулятором
type SwitchStatus struct {
	ID        string `json:"id" binding:"required"`
	State     bool   `json:"state" binding:"required"`
	Timestamp int64  `json:"timestamp" binding:"required"`
}
