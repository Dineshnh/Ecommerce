package models

import (
	"time"
)

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Category    string    `gorm:"index;not null" json:"category"`
	Price       float64   `gorm:"not null" json:"price"`
	Stock       int       `gorm:"default:0" json:"stock"`
	ImageURL    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
