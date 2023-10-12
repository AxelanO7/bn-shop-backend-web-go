package model

import (
	"gorm.io/gorm"
)

// Supplier struct
type Supplier struct {
	gorm.Model
	NameSupplier string `json:"name_supplier"`
	Phone        int    `json:"phone"`
	Address      string `json:"address"`
}

// Suppliers struct
type Suppliers struct {
	Suppliers []Supplier `json:"suppliers"`
}
