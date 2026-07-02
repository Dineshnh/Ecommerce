package controllers

import (
	"ecommerce-backend/config"
	"ecommerce-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminStats struct {
	Revenue float64 `json:"revenue"`
	Users   int64   `json:"users"`
	Orders  int64   `json:"orders"`
}

func GetAdminStats(c *gin.Context) {

	var stats AdminStats

	if err := config.DB.Model(&models.User{}).Count(&stats.Users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to count users"})
		return
	}

	if err := config.DB.Model(&models.Order{}).Count(&stats.Orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to count orders"})
		return
	}

	if err := config.DB.Model(&models.Order{}).
		Select("COALESCE(SUM(total),0)").
		Scan(&stats.Revenue).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to calculate revenue"})
		return
	}

	c.JSON(http.StatusOK, stats)
}
