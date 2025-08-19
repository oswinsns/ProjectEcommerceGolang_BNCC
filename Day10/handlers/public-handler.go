package handlers

import (
	"Day10/configs"
	"Day10/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func ShowProductsPage(c *gin.Context) {
	// Available products (example: stock > 0)
	var available []models.Product
	if err := configs.DB.Where("stock > ?", 0).Find(&available).Error; err != nil {
		c.String(http.StatusInternalServerError, "Error fetching available products")
		return
	}

	// Newest products (order by created_at desc, limit 5)
	var newest []models.Product
	if err := configs.DB.Order("created_at desc").Limit(5).Find(&newest).Error; err != nil {
		c.String(http.StatusInternalServerError, "Error fetching newest products")
		return
	}

	// Pass both lists to template
	c.HTML(http.StatusOK, "public.html", gin.H{
		"title":     "Products",
		"available": available,
		"newest":    newest,
	})
}

func ExportProducts(c *gin.Context) {
	var products []models.Product
	if err := configs.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create new Excel file
	f := excelize.NewFile()
	sheet := "Products"
	index, _ := f.NewSheet(sheet)

	// Headers
	headers := []string{"ID", "Name", "Price", "Stock", "Created At"}
	for i, h := range headers {
		col := string(rune('A' + i))
		f.SetCellValue(sheet, col+"1", h)
	}

	// Fill data
	for i, p := range products {
		row := strconv.Itoa(i + 2) // start at row 2
		f.SetCellValue(sheet, "A"+row, p.ID)
		f.SetCellValue(sheet, "B"+row, p.Name)
		f.SetCellValue(sheet, "C"+row, p.Price)
		f.SetCellValue(sheet, "D"+row, p.Stock)
		f.SetCellValue(sheet, "E"+row, p.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	// Set active sheet
	f.SetActiveSheet(index)

	// Write file directly to response
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", `attachment; filename="products.xlsx"`)
	c.Header("Content-Transfer-Encoding", "binary")

	if err := f.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
