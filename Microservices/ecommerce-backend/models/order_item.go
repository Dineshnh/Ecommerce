package models

type OrderItem struct {
	ID        uint `gorm:"primaryKey"`
	OrderID   uint
	ProductID uint
	Quantity  int     `gorm:"not null"`
	Price     float64 `gorm:"not null"`

	Order   Order
	Product Product
}
