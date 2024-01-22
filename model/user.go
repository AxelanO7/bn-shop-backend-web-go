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
	Username string `json:"username"`
	Password string `json:"password"`
	Status   int    `json:"status"`
}

// Users struct
type Users struct {
	Users []User `json:"users"`
}
