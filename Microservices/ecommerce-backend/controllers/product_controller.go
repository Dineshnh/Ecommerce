package controllers

import (
	"ecommerce-backend/config"
	"ecommerce-backend/models"
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := config.DB.Create(&product)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, product)
}

func GetProducts(c *gin.Context) {
	var products []models.Product

	result := config.DB.Find(&products)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	fmt.Println("Products:", products)
	c.JSON(200, products)
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")

	var product models.Product
	result := config.DB.First(&product, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(200, product)
}

func UpdateProductByID(c *gin.Context) {
	id := c.Param("id")

	var product models.Product
	result := config.DB.First(&product, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(200, product)
}

func DeleteProductByID(c *gin.Context) {
	id := c.Param("id")

	var product models.Product
	result := config.DB.First(&product, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(200, product)
}

func ProductSearch(c *gin.Context) {
	id := c.Param("id")

	var product models.Product
	result := config.DB.First(&product, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(200, product)
}

func ProductFilter(c *gin.Context) {
	id := c.Param("id")

	var product models.Product
	result := config.DB.First(&product, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(200, product)
}
