package model

import (
	"gorm.io/gorm"
)

// DetailOutput struct
type DetailOutput struct {
	gorm.Model
	IdOutput           int    `json:"id_output"`
	Output             Output `gorm:"foreignKey:OutputID" json:"output"`
	NameProductOutput  string `json:"name_product_output"`
	UnitProduct        string `json:"unit_product"`
	TotalProductOutput int    `json:"total_product_output"`
	TypeProduct        string `json:"type_product"`
	PriceUnit          int    `json:"price_unit"`
	TotalPrice         int    `json:"total_price"`
}

// DetailOutputs struct
type DetailOutputs struct {
	DetailOutputs []DetailOutput `json:"detail_outputs"`
}
