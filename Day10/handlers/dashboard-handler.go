package handlers

import (
	"Day10/configs"
	"Day10/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminDashboard(c *gin.Context) {
	var products []models.Product
	var userCount int64
	var activeUsers int64
	var productCount int64
	var availableProducts int64

	// Query products
	if err := configs.DB.Find(&products).Error; err != nil {
		c.String(http.StatusInternalServerError, "Error fetching products")
		return
	}

	configs.DB.Model(&models.User{}).Count(&userCount)
	configs.DB.Model(&models.User{}).Where("is_active = ?", true).Count(&activeUsers)
	configs.DB.Model(&models.Product{}).Count(&productCount)
	configs.DB.Model(&models.Product{}).Where("stock > ?", 0).Count(&availableProducts)

	// Render template with both products + stats
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title":              "Welcome to admin's E-commerce Dashboard",
		"products":           products,
		"user_count":         userCount,
		"active_users":       activeUsers,
		"total_products":     productCount,
		"available_products": availableProducts,
	})
}
