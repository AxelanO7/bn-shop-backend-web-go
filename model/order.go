package model

import (
	"gorm.io/gorm"
)

// Order struct
type Order struct {
	gorm.Model
	PurchaseOrder   string   `json:"purchase_order" gorm:"uniqueIndex:idx_purchase_order"`
	DateTransaction string   `json:"date_transaction"`
	IdSupplier      int      `json:"id_supplier"`
	Supplier        Supplier `gorm:"foreignKey:IdSupplier" json:"supplier"`
	TypeTransaction string   `json:"type_transaction"`
	Status          int      `json:"status"`
	IsConfirm       bool     `json:"is_confirm"`
}

// Orders struct
type Orders struct {
	Orders []Order `json:"orders"`
}
