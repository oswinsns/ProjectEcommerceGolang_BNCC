package handlers

import (
	"html/template"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/home.html")
	if err != nil {
		http.Error(w, "Error parsing template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	temp.Execute(w, nil)
}

// func Welcome(c *gin.Context) {
// 	temp, err := template.ParseFiles("views/home.html")
// 	if err != nil {
// 		c.String(http.StatusInternalServerError, "Error parsing template: %v", err)
// 		return
// 	}
// 	// c.Header("Content-Type", "text/html")
// 	// c.Writer.WriteHeader(http.StatusOK)
// 	temp.Execute(c.Writer, nil)
// }

// func Welcome(c *gin.Context) {
// 	c.HTML(http.StatusOK, "home.html", gin.H{
// 		"title": "Welcome to Day10 API",
// 	})
// }

// func(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	w.WriteHeader(http.StatusOK)
// 	_, err := w.Write([]byte("<h1>Welcome to Day10 API</h1>"))
// 	if err != nil {
// 		http.Error(w, "Error writing response", http.StatusInternalServerError)
// 	}
// }
