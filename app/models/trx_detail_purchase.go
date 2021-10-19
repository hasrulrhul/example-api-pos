package models

import (
	"time"

	"gorm.io/gorm"
)

type TrxDetailPurchase struct {
	ID         uint           `json:"id"`
	PurchaseID int64          `gorm:"not null" json:"purchase_id"`
	OutletID   int64          `gorm:"not null" json:"outlet_id"`
	ProductID  int64          `gorm:"not null" json:"product_id"`
	Qty        int64          `json:"qty"`
	Price      int64          `json:"price"`
	TotalPrice int64          `json:"total_price"`
	Product    Product        `gorm:"foreignKey:ProductID" json:"product"`
	Outlet     Outlet         `gorm:"foreignKey:OutletID" json:"outlet"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
