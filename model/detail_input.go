package model

import (
	"gorm.io/gorm"
)

// DetailInput struct
type DetailInput struct {
	gorm.Model
	IdInput     int    `json:"id_input"`
	Input       Input  `gorm:"foreignKey:IdInput" json:"input"`
	CodeProduct string `json:"code_product"`
	NameRaw     string `json:"name_raw"`
	UnitProduct string `json:"unit_product"`
	TotalUsed   int    `json:"total_used"`
	TypeProduct string `json:"type_product"`
	PriceUnit   int    `json:"price_unit"`
}

// DetailInputs struct
type DetailInputs struct {
	DetailInputs []DetailInput `json:"detail_inputs"`
}
