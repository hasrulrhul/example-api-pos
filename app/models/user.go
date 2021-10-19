package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Username  string         `gorm:"unique" json:"username"`
	Email     string         `gorm:"unique" json:"email"`
	Password  string         `json:"password"`
	Role      string         `gorm:"type:enum('superadmin', 'admin');default:'admin'" json:"role"`
	Type      string         `gorm:"type:enum('merchant', 'outlet');default:'outlet'" json:"type"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type LoginCredentials struct {
	// Username string `gorm:"unique" json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}
