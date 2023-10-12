package model

import (
	"gorm.io/gorm"
)

// Input struct
type Input struct {
	gorm.Model
	DateInput       string `json:"date_input"`
	CodeProduct     int    `json:"code_product"`
	NameProduct     string `json:"name_product"`
	TypeProduct     string `json:"type_product"`
	TotalProduction int    `json:"total_production"`
	PriceProduct    int    `json:"price_product"`
}

// Inputs struct
type Inputs struct {
	Inputs []Input `json:"inputs"`
}
