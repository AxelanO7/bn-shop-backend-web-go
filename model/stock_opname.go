
package model

import (
	"gorm.io/gorm"
)

// StockOpname struct
type StockOpname struct {
	gorm.Model
	DateCalculate string `json:"date_calculate"`
	NameProduct   string `json:"name_product"`
	StockReal     int    `json:"stock_real"`
}

// StockOpnames struct
type StockOpnames struct {
	StockOpnames []StockOpname `json:"stockopnames"`
}
