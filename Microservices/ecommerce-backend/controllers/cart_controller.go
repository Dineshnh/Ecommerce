package controllers

import (
	"net/http"
	"strconv"

	"ecommerce-backend/config"
	"ecommerce-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddToCart(c *gin.Context) {

	userID := c.MustGet("userID").(uint)

	var input struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Quantity must be greater than zero",
		})
		return
	}

	var product models.Product
	if err := config.DB.First(&product, input.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product not found",
		})
		return
	}

	var cart models.Cart
	config.DB.Where("user_id = ?", userID).FirstOrCreate(&cart, models.Cart{
		UserID: userID,
	})

	var item models.CartItem
	err := config.DB.Where("cart_id = ? AND product_id = ?", cart.ID, input.ProductID).
		First(&item).Error

	switch err {
	case nil:
		item.Quantity += input.Quantity
		config.DB.Save(&item)
	case gorm.ErrRecordNotFound:
		item = models.CartItem{
			CartID:    cart.ID,
			ProductID: input.ProductID,
			Quantity:  input.Quantity,
		}
		config.DB.Create(&item)
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to add product to cart",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product added to cart successfully",
	})
}

func GetCart(c *gin.Context) {

	userID := c.MustGet("userID").(uint)

	var cart models.Cart

	if err := config.DB.Where("user_id = ?", userID).First(&cart).Error; err != nil {

		c.JSON(http.StatusOK, gin.H{
			"cart_id": nil,
			"items":   []models.CartItem{},
		})
		return
	}

	var items []models.CartItem

	config.DB.Preload("Product").
		Where("cart_id = ?", cart.ID).
		Find(&items)

	c.JSON(http.StatusOK, gin.H{
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Quantity must be greater than zero",
		})
		return
	}

	var cart models.Cart

	if err := config.DB.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Cart not found",
		})
		return
	}

	var item models.CartItem

	if err := config.DB.Where("cart_id = ? AND product_id = ?", cart.ID, input.ProductID).
		First(&item).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Item not found in cart",
		})
		return
	}

	item.Quantity = input.Quantity
	config.DB.Save(&item)

	c.JSON(http.StatusOK, gin.H{
		"message": "Cart updated successfully",
	})
}

func RemoveCartItem(c *gin.Context) {

	userID := c.MustGet("userID").(uint)

	productIDStr := c.Query("product_id")

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid product_id",
		})
		return
	}

	var cart models.Cart

	if err := config.DB.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Cart not found",
		})
		return
	}

	result := config.DB.Where(
		"cart_id = ? AND product_id = ?",
		cart.ID,
		uint(productID),
	).Delete(&models.CartItem{})

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Item not found in cart",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Item removed from cart successfully",
	})
}
