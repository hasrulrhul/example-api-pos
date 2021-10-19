package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductOutlet struct {
	ID           uint    `json:"id"`
	ProductID    int64   `gorm:"not null" json:"product_id"`
	OutletID     int64   `gorm:"not null" json:"outlet_id"`
	Qty          int64   `json:"qty"`
	SellingPrice int64   `json:"selling_price"`
	Product      Product `gorm:"foreignKey:ProductID" json:"product"`
	Outlet       Outlet  `gorm:"foreignKey:OutletID" json:"outlet"`
	// DetailPurchase TrxDetailPurchase `gorm:"foreignKey:PurchaseID" json:"detailpurchase"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
