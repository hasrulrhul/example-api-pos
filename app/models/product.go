package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `json:"id"`
	ProductName string         `json:"product_name"`
	Sku         string         `gorm:"unique" json:"sku"`
	Category    string         `json:"category"`
	Price       int64          `json:"price"`
	Image       string         `json:"image"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
