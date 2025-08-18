package routes

import (
	"Day10/handlers"
	"Day10/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Public API
	r.GET("/products/latest", handlers.GetLatestProducts)
	r.GET("/products/available", handlers.GetAvailableProducts)
	r.GET("/", handlers.AdminDashboard)
	r.POST("/admin/login", handlers.Login)
	r.GET("/login", handlers.ShowLogin)
	r.POST("/login", handlers.Login)

	// Admin API (protected)
	admin := r.Group("/admin", middlewares.AuthMiddleware())
	{
		// TODO: add CRUD users & products, export endpoint

		admin.GET("/dashboard", handlers.AdminDashboard)
		admin.GET("/products/export", handlers.ExportProducts)

		// User CRUD
		admin.GET("/users", handlers.GetUsers)
		admin.GET("/users/create", handlers.CreateUsers)
		admin.POST("/add-users", handlers.CreateUser)
		admin.GET("/users/edit/:id", handlers.EditUserForm)
		admin.POST("/users/update/:id", handlers.UpdateUser)
		admin.POST("/users/delete/:id", handlers.DeleteUser)

		// Product CRUD
		// admin.GET("/products", handlers.GetProducts)
		admin.POST("/add-products", handlers.CreateProducts)
		admin.GET("/products/create", handlers.CreateProduct)
		admin.PUT("/products/:id", handlers.UpdateProduct)
		admin.DELETE("/products/:id", handlers.DeleteProduct)
	}
}
