package model

import (
	"gorm.io/gorm"
)

// Output struct
type Output struct {
	gorm.Model
	NoOutput   string `json:"no_output"`
	DateOutput string `json:"date_output"`
	IdUser     int    `json:"id_user"`
	User       User   `gorm:"foreignKey:IdUser" json:"user"`
}

// Outputs struct
type Outputs struct {
	Outputs []Output `json:"outputs"`
}
