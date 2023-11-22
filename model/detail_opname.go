package model

import (
	"gorm.io/gorm"
)

// DetailOpname struct
type DetailOpname struct {
	gorm.Model
	IdOpname     int         `json:"id_opname"`
	Opname       StockOpname `gorm:"foreignKey:IdOpname" json:"opname"`
	CodeProduct  string      `json:"code_product"`
	NameFinished string      `json:"name_finished"`
	UnitProduct  string      `json:"unit_product"`
	TypeProduct  string      `json:"type_product"`
	PriceUnit    int         `json:"price_unit"`
	StockSystem  int         `json:"stock_system"`
	StockReal    int         `json:"stock_real"`
	TotalDiff    int         `json:"total_diff"`
}

// DetailOpnames struct
type DetailOpnames struct {
	DetailOpnames []DetailOpname `json:"detail_opnames"`
}
