package model

import (
	"gorm.io/gorm"
)

// Stock struct
type Stock struct {
	gorm.Model
	CodeProduct  string   `json:"code_product" gorm:"uniqueIndex:idx_code_product"`
	NameProduct  string   `json:"name_product"`
	UnitProduct  string   `json:"unit_product"`
	TotalProduct float64  `json:"total_product"`
	TypeProduct  string   `json:"type_product"`
	PriceProduct int      `json:"price_product"`
	IdSupplier   int      `json:"id_supplier"`
	Supplier     Supplier `gorm:"foreignKey:IdSupplier" json:"supplier"`
	IdUser       int      `json:"id_user"`
	User         User     `gorm:"foreignKey:IdUser" json:"user"`
}

// Stocks struct
type Stocks struct {
	Stocks []Stock `json:"stocks"`
}
