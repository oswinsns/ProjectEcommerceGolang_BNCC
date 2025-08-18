package main

import (
	"Day10/configs"
	"Day10/databases"
	"Day10/databases/seeders"

	// "Day10/handlers"
	"fmt"
	// "log"

	// "Day10/databases/seeders"
	"Day10/routes"

	// "net/http"

	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi Kesalahan:", r)
		}
	}()

	// validasi env
	if err := godotenv.Load(); err != nil {
		panic(fmt.Sprintf("No .env file found: %s", err.Error()))
	}

	// koneksi ke db
	configs.SetupMySQL()

	// automigrate
	databases.AutoMigrate()

	// seeder
	seeders.SeedProducts()
	seeders.SeedUsers()
	// seeders.SeederMataKuliah()
	// seeders.SeederPasswordDosen()

	// router
	r := gin.Default()

	// load HTML templates from views folder
	r.LoadHTMLGlob("views/*")

	r.Static("/static", "./Day10/static")

	// route: serve homepage

	r.GET("/admin/login", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{
			"title": "Welcome to Day10 API",
		})
	})

	r.GET("/admin/user", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{
			"title": "Welcome to Day10 API",
		})
	})

	// register custom funcs
	r.SetFuncMap(template.FuncMap{
		"add": func(a, b int) int { return a + b },
	})

	routes.SetupRoutes(r)

	r.Run()
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Welcome to Day10 API")
	// }
}
