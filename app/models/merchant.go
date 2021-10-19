package models

import (
	"time"

	"gorm.io/gorm"
)

type Merchant struct {
	ID           uint           `json:"id"`
	MerchantName string         `gorm:"unique" json:"merchant_name"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
