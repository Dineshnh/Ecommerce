package models

type OrderItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`

	Order   Order   `gorm:"foreignKey:OrderID" json:"-"`
	Product Product `gorm:"foreignKey:ProductID" json:"product"`
}
