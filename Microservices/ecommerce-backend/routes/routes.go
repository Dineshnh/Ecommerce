package routes

import (
	"ecommerce-backend/controllers"
	"ecommerce-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	public := r.Group("/")
	{
		public.GET("/products", controllers.GetProducts)
		public.GET("/products/:id", controllers.GetProductByID)
		public.GET("/products/filter", controllers.ProductFilter)
		public.GET("/products/search", controllers.ProductSearch)
	}

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", controllers.GetProfile)
		protected.PUT("/profile", controllers.UpdateProfile)

		protected.POST("/products", controllers.CreateProduct)
		protected.PUT("/products/:id", controllers.UpdateProductByID)
		protected.DELETE("/products/:id", controllers.DeleteProductByID)
	}

	cart := r.Group("/cart")
	cart.Use(middleware.AuthMiddleware())
	{
		cart.POST("/add", controllers.AddToCart)
		cart.GET("", controllers.GetCart)
		cart.PUT("/update", controllers.UpdateCartItem)
		cart.DELETE("/remove", controllers.RemoveCartItem)
	}

	order := r.Group("/orders")
	order.Use(middleware.AuthMiddleware())
	{
		order.POST("/place", controllers.PlaceOrder)
		order.GET("", controllers.GetMyOrders)
		order.GET("/:id", controllers.GetOrderDetails)
	}

	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.AdminMiddleware())
	{
		admin.POST("/products", controllers.CreateProduct)
		admin.PUT("/products/:id", controllers.UpdateProductByID)
		admin.DELETE("/products/:id", controllers.DeleteProductByID)
	}
}
