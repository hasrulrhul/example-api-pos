package models

import (
	"time"

	"gorm.io/gorm"
)

type TrxDetailSale struct {
	ID         uint           `json:"id"`
	SaleID     int64          `json:"sale_id"`
	OutletID   int64          `json:"outlet_id"`
	ProductID  int64          `json:"product_id"`
	Qty        int64          `json:"qty"`
	Price      int64          `json:"price"`
	TotalPrice int64          `json:"total_price"`
	Product    Product        `gorm:"foreignKey:ProductID" json:"product"`
	Outlet     Outlet         `gorm:"foreignKey:OutletID" json:"outlet"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
