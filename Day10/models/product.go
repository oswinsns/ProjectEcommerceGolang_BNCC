package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"unique;primaryKey"`
	Name      string    `json:"name" form:"name"`
	Stock     int       `json:"stock" form:"stock,min=0"`
	Price     float64   `json:"price" form:"price,min=0"`
	CreatedAt time.Time `json:"created_at"`
}
