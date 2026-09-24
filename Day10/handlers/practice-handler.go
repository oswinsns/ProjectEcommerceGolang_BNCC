package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PracticeLab merender halaman latihan interaktif berbasis web
func PracticeLab(c *gin.Context) {
	c.HTML(http.StatusOK, "practice.html", gin.H{
		"title": "🎮 E-Commerce Coding Practice Lab",
	})
}
