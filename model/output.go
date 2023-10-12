package model

import (
	"gorm.io/gorm"
)

// Output struct
type Output struct {
	gorm.Model
	DateOutput string `json:"date_output"`
}

// Outputs struct
type Outputs struct {
	Outputs []Output `json:"outputs"`
}
