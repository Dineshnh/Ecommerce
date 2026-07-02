package config

import (
	"fmt"
	"log"

	"ecommerce-backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := "root:root@tcp(127.0.0.1:3306)/ecommerce_db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("❌ Failed to connect database:", err)
		return
	}

	DB = db

	err = DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
	)

	if err != nil {
		log.Fatal("❌ AutoMigrate failed:", err)
		return
	}

	fmt.Println("✅ Database connected & migrated successfully")
}
