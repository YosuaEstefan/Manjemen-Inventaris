package models

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	ProductID uint   `gorm:"primaryKey" json:"id"`
	Location  string `gorm:"primaryKey" json:"location"`
	Quantity  int    `json:"quantity"`
}
