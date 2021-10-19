package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID          uint           `json:"id"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	Address     string         `json:"address"`
	Email       string         `gorm:"unique" json:"email"`
	PhoneNumber string         `gorm:"unique" json:"phone_number"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
