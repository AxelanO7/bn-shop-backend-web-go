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
	PriceUnit    float64     `json:"price_unit"`
	StockSystem  float64     `json:"stock_system"`
	StockReal    float64     `json:"stock_real"`
	TotalDiff    float64     `json:"total_diff"`
}

// DetailOpnames struct
type DetailOpnames struct {
	DetailOpnames []DetailOpname `json:"detail_opnames"`
}
