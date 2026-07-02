package main

import (
	"ecommerce-backend/config"
	"ecommerce-backend/routes"

	"github.com/gin-gonic/gin"
)

func main(){
	config.ConnectDB()
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}