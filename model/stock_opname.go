package model

import (
	"gorm.io/gorm"
)

// StockOpname struct
type StockOpname struct {
	gorm.Model
	DateCalculate   string `json:"date_calculate"`
	CodeStockOpname string `json:"code_stock_opname"`
}

// StockOpnames struct
type StockOpnames struct {
	StockOpnames []StockOpname `json:"stockopnames"`
}
