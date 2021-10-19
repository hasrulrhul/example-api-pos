package models

import (
	"time"

	"gorm.io/gorm"
)

type CartSale struct {
	ID         uint           `json:"id"`
	CustomerID int64          `gorm:"not null" json:"customer_id"`
	OutletID   int64          `gorm:"not null" json:"outlet_id"`
	ProductID  int64          `gorm:"not null" json:"product_id"`
	Qty        int64          `json:"qty"`
	Price      int64          `json:"price"`
	TotalPrice int64          `json:"total_price"`
	Customer   Customer       `gorm:"foreignKey:CustomerID" json:"customer"`
	Outlet     Outlet         `gorm:"foreignKey:OutletID" json:"outlet"`
	Product    Product        `gorm:"foreignKey:ProductID" json:"product"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
