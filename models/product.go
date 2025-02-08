package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID           uint        `gorm:"primaryKey" json:"id"`
	Name_Product string      `json:"Name_Product "`
	Description  string      `json:"description"`
	Price        float64     `json:"price"`
	Category     string      `json:"category"`
	Inventories  []Inventory `gorm:"foreignKey:ID "`
}
