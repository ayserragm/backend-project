package models

import "time"

type AuditLog struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	EntityType string    `json:"entity_type"`
	EntityID   uint      `json:"entity_id"`
	Action     string    `json:"action"`
	Details    string    `json:"details"`
	CreatedAt  time.Time `json:"created_at"`
}
