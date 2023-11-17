package model

import (
	"gorm.io/gorm"
)

// General Cart struct
type DetailOrder struct {
	gorm.Model
	CodeProduct  string  `json:"code_product" gorm:"uniqueIndex:idx_code_product"`
	IdOrder      int     `json:"id_order"`
	Order        Order   `gorm:"foreignKey:IdOrder" json:"order"`
	NameProduct  string  `json:"name_product"`
	UnitProduct  string  `json:"unit_product"`
	TypeProduct  string  `json:"type_product"`
	PriceProduct int     `json:"price_product"`
	TotalOrder   float64 `json:"total_order"`
}

// General Carts struct
type DetailOrders struct {
	DetailJournals []DetailOrder `json:"detail_journals"`
}
