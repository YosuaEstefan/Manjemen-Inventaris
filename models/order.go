package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProductID uint      `json:"product_id"`
	Quantity  int       `json:"quantity"`
	OrderDate time.Time `json:"order_date"`
}
