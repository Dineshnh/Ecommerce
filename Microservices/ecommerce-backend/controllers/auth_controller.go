package controllers

import (
	"fmt"
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

	fmt.Println("Original Password:", user.Password)

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Password hashing failed",
		})
		return
	}
	fmt.Println("Generated Hash:", hashedPassword)

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

// func Login(c *gin.Context) {
// 	var input struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}

// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error":  "Invalid JSON input",
// 			"detail": err.Error(),
// 		})
// 		return
// 	}

// 	var user models.User

// 	result := config.DB.Where("email = ?", input.Email).First(&user)
// 	if result.Error != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"message": "Invalid Email or Password",
// 		})
// 		return
// 	}

// 	if !utils.CheckPasswordHash(input.Password, user.Password) {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"message": "Invalid Email or Password",
// 		})
// 		return
// 	}

// 	token, err := utils.GenerateToken(user.ID, user.Role)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "Failed to generate token",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Login Successful",
// 		"token":   token,
// 		"user": gin.H{
// 			"id":   user.ID,
// 			"name": user.Name,
// 			"role": user.Role,
// 		},
// 	})
// }

func Login(c *gin.Context) {

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// 🔍 DEBUG: print raw request
	fmt.Println("=== LOGIN REQUEST ===")
	fmt.Println("Email input:", input.Email)

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("JSON ERROR:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid JSON input",
			"detail": err.Error(),
		})
		return
	}

	fmt.Println("Parsed Email:", input.Email)
	fmt.Println("Parsed Password:", input.Password)

	var user models.User

	result := config.DB.Where("email = ?", input.Email).First(&user)

	// 🔍 DEBUG DB RESULT
	if result.Error != nil {
		fmt.Println("DB ERROR:", result.Error)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid Email or Password",
		})
		return
	}

	fmt.Println("USER FOUND IN DB")
	fmt.Println("DB Email:", user.Email)
	fmt.Println("DB Password (HASH):", user.Password)
	fmt.Println("DB Role:", user.Role)

	// 🔍 PASSWORD CHECK DEBUG
	match := utils.CheckPasswordHash(input.Password, user.Password)
	fmt.Println("PASSWORD MATCH RESULT:", match)

	if !match {
		fmt.Println("PASSWORD INCORRECT")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid Email or Password",
		})
		return
	}

	// token generation
	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		fmt.Println("TOKEN ERROR:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	fmt.Println("LOGIN SUCCESS for user:", user.Email)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successful",
		"token":   token,
		"user": gin.H{
			"id":   user.ID,
			"name": user.Name,
			"role": user.Role,
		},
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
