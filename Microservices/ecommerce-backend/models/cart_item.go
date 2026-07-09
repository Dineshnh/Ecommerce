package models

type CartItem struct {
	ID        uint `json:"id"`
	CartID    uint `json:"cart_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`

	Product Product `gorm:"foreignKey:ProductID" json:"product"`
}
