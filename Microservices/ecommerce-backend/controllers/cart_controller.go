package controllers

import (
	"ecommerce-backend/config"
	"ecommerce-backend/models"

	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var input struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var cart models.Cart
	config.DB.Where("user_id = ?", userID).FirstOrCreate(&cart, models.Cart{
		UserID: userID,
	})

	var item models.CartItem
	err := config.DB.Where("cart_id = ? AND product_id = ?", cart.ID, input.ProductID).
		First(&item).Error

	if err == nil {
		item.Quantity += input.Quantity
		config.DB.Save(&item)
	} else {
		item = models.CartItem{
			CartID:    cart.ID,
			ProductID: input.ProductID,
			Quantity:  input.Quantity,
		}
		config.DB.Create(&item)
	}

	c.JSON(200, gin.H{
		"message": "Added to cart",
	})
}


func GetCart(c *gin.Context) {

	userID := c.MustGet("userID").(uint)

	var cart models.Cart
	config.DB.Where("user_id = ?", userID).First(&cart)

	var items []models.CartItem
	config.DB.Where("cart_id = ?", cart.ID).Find(&items)

	c.JSON(200, gin.H{
		"cart_id": cart.ID,
		"items":   items,
	})
}


func UpdateCartItem(c *gin.Context) {

	userID := c.MustGet("userID").(uint)

	var input struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var cart models.Cart
	config.DB.Where("user_id = ?", userID).First(&cart)

	var item models.CartItem
	result := config.DB.Where("cart_id = ? AND product_id = ?", cart.ID, input.ProductID).
		First(&item)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}

	item.Quantity = input.Quantity
	config.DB.Save(&item)

	c.JSON(200, gin.H{"message": "Cart updated"})
}



func RemoveCartItem(c *gin.Context) {

	userID := c.MustGet("userID").(uint)

	productID := c.Query("product_id")

	var cart models.Cart
	config.DB.Where("user_id = ?", userID).First(&cart)

	config.DB.Where("cart_id = ? AND product_id = ?", cart.ID, productID).
		Delete(&models.CartItem{})

	c.JSON(200, gin.H{
		"message": "Item removed",
	})
}
