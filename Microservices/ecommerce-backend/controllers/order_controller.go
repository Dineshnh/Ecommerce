package controllers

import (
	"ecommerce-backend/config"
	"ecommerce-backend/models"

	"github.com/gin-gonic/gin"
)

func PlaceOrder(c *gin.Context) {

	userID := c.MustGet("userID").(uint)

	var cart models.Cart
	config.DB.Where("user_id = ?", userID).First(&cart)

	var cartItems []models.CartItem
	config.DB.Where("cart_id = ?", cart.ID).Find(&cartItems)

	if len(cartItems) == 0 {
		c.JSON(400, gin.H{"error": "Cart is empty"})
		return
	}

	var total float64 = 0

	order := models.Order{
		UserID: userID,
		Status: "pending",
	}

	config.DB.Create(&order)

	for _, item := range cartItems {

		var product models.Product
		config.DB.First(&product, item.ProductID)

		total += product.Price * float64(item.Quantity)

		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		}

		config.DB.Create(&orderItem)
	}

	order.Total = total
	config.DB.Save(&order)

	config.DB.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{})

	c.JSON(200, gin.H{
		"message":  "Order placed successfully",
		"order_id": order.ID,
		"total":    total,
	})
}

func GetMyOrders(c *gin.Context) {

	userID := c.MustGet("userID").(uint)

	var orders []models.Order
	config.DB.Where("user_id = ?", userID).Find(&orders)

	c.JSON(200, gin.H{
		"orders": orders,
	})
}

func GetOrderDetails(c *gin.Context) {

	orderID := c.Param("id")

	var order models.Order
	config.DB.First(&order, orderID)

	var items []models.OrderItem
	config.DB.Where("order_id = ?", orderID).Find(&items)

	c.JSON(200, gin.H{
		"order": order,
		"items": items,
	})
}


