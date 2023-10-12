package model

import (
	"gorm.io/gorm"
)

// Login struct
type Login struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

// Logins struct
type Logins struct {
	Logins []Login `json:"logins"`
}
