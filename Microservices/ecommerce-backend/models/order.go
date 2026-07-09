package models

import "time"

type Order struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	Total     float64   `json:"total"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`

	User       User        `gorm:"foreignKey:UserID" json:"-"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID" json:"order_items,omitempty"`
}
