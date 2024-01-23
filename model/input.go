package model

import (
	"gorm.io/gorm"
)

// Input struct
type Input struct {
	gorm.Model
	NoInput      string `json:"no_input"`
	DateInput    string `json:"date_input"`
	CodeProduct  string `json:"code_product"`
	NameProduct  string `json:"name_product"`
	TypeProduct  string `json:"type_product"`
	TotalProduct int    `json:"total_product"`
	PriceProduct int    `json:"price_product"`
	IdUser       int    `json:"id_user"`
	User         User   `gorm:"foreignKey:IdUser" json:"user"`
}

// Inputs struct
type Inputs struct {
	Inputs []Input `json:"inputs"`
}
