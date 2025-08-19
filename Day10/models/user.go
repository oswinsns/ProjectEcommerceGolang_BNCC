package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// Nama      string `json:"nama" binding:"required"`
	// NIK       string `json:"nik" binding:"required"`
	// Password  string `json:"password" binding:"required"`

	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	IsActive bool   `json:"is_active" form:"is_active"`
	Email    string `json:"email" form:"email"`
}
