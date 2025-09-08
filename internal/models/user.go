package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Username     string    `gorm:"uniqueIndex;size:190" json:"username"`
	Email        string    `gorm:"uniqueIndex;size:190" json:"email"`
	PasswordHash string    `json:"-"`
	Role         string    `gorm:"size:32" json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
