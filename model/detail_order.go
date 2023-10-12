package model

import (
	"gorm.io/gorm"
)

// General Cart struct
type DetailOrder struct {
	gorm.Model
	IdOrder      int    `json:"id_order"`
	Order        Order  `gorm:"foreignKey:IdOrder" json:"order"`
	NameProduct  string `json:"name_product"`
	UnitProduct  string `json:"unit_product"`
	TypeProduct  string `json:"type_product"`
	PriceProduct int    `json:"price_product"`
	TotalOrder   int    `json:"total_order"`
}

// General Carts struct
type DetailOrders struct {
	DetailJournals []DetailOrder `json:"detail_journals"`
}
