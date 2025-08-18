package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"unique;primaryKey"`
	Name      string    `json:"name" binding:"required"`
	Stock     int       `json:"stock" binding:"required,min=0"`
	Price     float64   `json:"price" binding:"required,min=0"`
	CreatedAt time.Time `json:"created_at"`
}
