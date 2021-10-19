package models

import (
	"time"

	"gorm.io/gorm"
)

type Outlet struct {
	ID         uint           `json:"id"`
	MerchantID int64          `gorm:"not null" json:"merchant_id"`
	OutletName string         `gorm:"unique" json:"outlet_name"`
	Merchant   Merchant       `gorm:"foreignKey:MerchantID" json:"merchant"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
