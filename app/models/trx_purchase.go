package models

import (
	"time"

	"gorm.io/gorm"
)

type TrxPurchase struct {
	ID           uint                `json:"id"`
	OutletID     int64               `gorm:"not null" json:"outlet_id"`
	TotalPayment int64               `json:"total_payment"`
	Status       string              `gorm:"type:enum('paid', 'unpaid');default:'unpaid'" json:"status"`
	Outlet       Outlet              `gorm:"foreignKey:OutletID" json:"outlet"`
	Detail       []TrxDetailPurchase `gorm:"foreignKey:PurchaseID" json:"detail"`
	CreatedAt    time.Time           `json:"created_at"`
	UpdatedAt    time.Time           `json:"updated_at"`
	DeletedAt    gorm.DeletedAt      `gorm:"index" json:"deleted_at"`
}
