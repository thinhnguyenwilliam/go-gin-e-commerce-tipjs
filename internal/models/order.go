package model

import (
	"time"
)

type Order struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `gorm:"not null"`
	User      User    `gorm:"foreignKey:UserID"`
	Total     float64 `gorm:"not null"`
	Status    string  `gorm:"default:'pending'"` // pending, paid, shipped, cancelled
	CreatedAt time.Time
	UpdatedAt time.Time
	Items     []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID        uint    `gorm:"primaryKey"`
	OrderID   uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int     `gorm:"not null"`
	Price     float64 `gorm:"not null"` // snapshot of product price
}
