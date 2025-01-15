package entity

import (
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Name  string  `valid: "-"`
	Price float64 `valid: "required~"`
	SKU   string
}
