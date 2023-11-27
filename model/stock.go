package model

import (
	"gorm.io/gorm"
)

// Stock struct
type Stock struct {
	gorm.Model
	CodeProduct  string  `json:"code_product" gorm:"uniqueIndex:idx_code_product"`
	NameProduct  string  `json:"name_product"`
	UnitProduct  string  `json:"unit_product"`
	TotalProduct float64 `json:"total_product"`
	TypeProduct  string  `json:"type_product"`
	PriceProduct int     `json:"price_product"`
}

// Stocks struct
type Stocks struct {
	Stocks []Stock `json:"stocks"`
}
