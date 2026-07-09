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

	total := 0.0

	order := models.Order{
		UserID: userID,
		Status: "Pending",
	}

	config.DB.Create(&order)

	for _, item := range cartItems {

		var product models.Product

		// Check product exists
		if err := config.DB.First(&product, item.ProductID).Error; err != nil {
			c.JSON(404, gin.H{"error": "Product not found"})
			return
		}

		// Check stock
		if product.Stock < item.Quantity {
			c.JSON(400, gin.H{
				"error": product.Name + " is out of stock",
			})
			return
		}

		total += product.Price * float64(item.Quantity)

		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: product.ID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		}

		config.DB.Create(&orderItem)

		// Reduce stock
		product.Stock -= item.Quantity
		config.DB.Save(&product)
	}

	order.Total = total
	config.DB.Save(&order)

	// Clear cart
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

	if err := config.DB.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&orders).Error; err != nil {

		c.JSON(500, gin.H{
			"error": "Failed to fetch orders",
		})
		return
	}

	c.JSON(200, gin.H{
		"orders": orders,
	})
}

func GetOrderDetails(c *gin.Context) {

	userID := c.MustGet("userID").(uint)
	orderID := c.Param("id")

	var order models.Order

	// Ensure the order belongs to the logged-in user
	if err := config.DB.
		Where("id = ? AND user_id = ?", orderID, userID).
		First(&order).Error; err != nil {

		c.JSON(404, gin.H{
			"error": "Order not found",
		})
		return
	}

	var items []models.OrderItem

	if err := config.DB.
		Preload("Product").
		Where("order_id = ?", order.ID).
		Find(&items).Error; err != nil {

		c.JSON(500, gin.H{
			"error": "Failed to fetch order items",
		})
		return
	}

	c.JSON(200, gin.H{
		"order": order,
		"items": items,
	})
}
