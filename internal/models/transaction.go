package models

import "time"

type Transaction struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	FromUserID uint      `json:"from_user_id"`
	ToUserID   uint      `json:"to_user_id"`
	Amount     float64   `json:"amount"`
	Type       string    `json:"type"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}
