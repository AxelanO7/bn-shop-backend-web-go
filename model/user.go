package model

import (
	// "github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	Name     string `json:"name_user"`
	Position string `json:"position"`
}

// Users struct
type Users struct {
	Users []User `json:"users"`
}
