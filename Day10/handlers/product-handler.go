package handlers

import (
	"Day10/configs"
	"Day10/models"
	"net/http"

	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func GetLatestProducts(c *gin.Context) {
	var products []models.Product
	configs.DB.Order("created_at desc").Limit(5).Find(&products)
	c.JSON(http.StatusOK, products)
}

func GetAvailableProducts(c *gin.Context) {
	var products []models.Product
	configs.DB.Where("stock > ?", 0).Find(&products)
	c.JSON(http.StatusOK, products)
}

func ListProducts(c *gin.Context) {
	var products []models.Product
	if err := configs.DB.Find(&products).Error; err != nil {
		c.String(http.StatusInternalServerError, "Error fetching products")
		return
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"title":    "Latest Available Products",
		"products": products,
	})
}

// Create product
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	configs.DB.Create(&product)
	c.JSON(http.StatusCreated, product)
}

// Update product
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := configs.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	configs.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

// Delete product
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := configs.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

func ExportProducts(c *gin.Context) {
	var products []models.Product
	configs.DB.Find(&products)

	f := excelize.NewFile()
	sheet := "Products"
	index, _ := f.NewSheet(sheet)

	// Header
	headers := []string{"ID", "Name", "Stock", "CreatedAt"}
	for i, h := range headers {
		cell := fmt.Sprintf("%c1", 'A'+i)
		f.SetCellValue(sheet, cell, h)
	}

	// Rows
	for i, p := range products {
		row := i + 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), p.ID)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), p.Name)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), p.Stock)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), p.CreatedAt.Format(time.RFC3339))
	}

	f.SetActiveSheet(index)

	// Write file to response
	filename := "products.xlsx"
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("File-Name", filename)
	c.Header("Content-Transfer-Encoding", "binary")
	_ = f.Write(c.Writer)
}
