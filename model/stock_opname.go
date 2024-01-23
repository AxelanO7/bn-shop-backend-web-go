package model

import (
	"gorm.io/gorm"
)

// StockOpname struct
type StockOpname struct {
	gorm.Model
	DateCalculate   string   `json:"date_calculate"`
	CodeStockOpname string   `json:"code_stock_opname"`
	IdSupplier      int      `json:"id_supplier"`
	Supplier        Supplier `gorm:"foreignKey:IdSupplier" json:"supplier"`
	IdUser          int      `json:"id_user"`
	User            User     `gorm:"foreignKey:IdUser" json:"user"`
}

// StockOpnames struct
type StockOpnames struct {
	StockOpnames []StockOpname `json:"stockopnames"`
}
