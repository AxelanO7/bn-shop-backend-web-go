package model

import (
	"gorm.io/gorm"
)

// DetailOutput struct
type DetailOutput struct {
	gorm.Model
	IdOutput     int     `json:"id_output"`
	Output       Output  `gorm:"foreignKey:IdOutput" json:"output"`
	CodeProduct  string  `json:"code_product"`
	NameFinished string  `json:"name_finished"`
	UnitProduct  string  `json:"unit_product"`
	TotalUsed    float64 `json:"total_used"`
	TypeProduct  string  `json:"type_product"`
	PriceUnit    int     `json:"price_unit"`
}

// DetailOutputs struct
type DetailOutputs struct {
	DetailOutputs []DetailOutput `json:"detail_outputs"`
}
