package controllers

import (
	"net/http"

	"ecommerce-backend/config"
	"ecommerce-backend/models"
	"ecommerce-backend/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Password hashing failed",
		})
		return
	}

	user.Password = hashedPassword
	user.Role = "CUSTOMER"
	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Registered successfully",
	})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error":  "Invalid JSON input",
			"detail": err.Error(),
		})
		return
	}

	var user models.User

	result := config.DB.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		c.JSON(401, gin.H{"message": "Invalid Email or Password"})
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(401, gin.H{"message": "Invalid Email or Password"})
		return
	}

	token, _ := utils.GenerateToken(user.ID, user.Role)
	c.JSON(200, gin.H{
		"message": "Login Successful",
		"token":   token,
	})
}

func GetProfile(c *gin.Context) {

	userID := c.MustGet("userID").(uint)

	var user models.User

	result := config.DB.Select("id, name, email").First(&user, userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}

func UpdateProfile(c *gin.Context) {

	userID := c.MustGet("userID").(uint)

	var user models.User

	result := config.DB.First(&user, userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user.Name = input.Name
	user.Email = input.Email

	config.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully",
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
