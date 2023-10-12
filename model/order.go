package model

import (
	"gorm.io/gorm"
)

// Order struct
type Order struct {
	gorm.Model
	DateTransaction string   `json:"date_transaction"`
	IdSupplier      int      `json:"id_supplier"`
	Supplier        Supplier `gorm:"foreignKey:IdSupplier" json:"supplier"`
	TypeTransaction string   `json:"type_transaction"`
}

// Orders struct
type Orders struct {
	Orders []Order `json:"orders"`
}
