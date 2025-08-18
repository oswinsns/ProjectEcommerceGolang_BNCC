package handlers

import (
	"Day10/configs"
	"Day10/models"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

// List all users
func GetUsers(c *gin.Context) {
	var users []models.User
	if err := configs.DB.Find(&users).Error; err != nil {
		c.String(http.StatusInternalServerError, "Error fetching users")
		return
	}

	// Render HTML instead of JSON
	c.HTML(http.StatusOK, "users.html", gin.H{
		"title": "All Users",
		"users": users,
	})
}

// Create new user
// func CreateUser(c *gin.Context) {
// 	var user models.User
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	configs.DB.Create(&user)
// 	c.JSON(http.StatusCreated, user)
// }

func CreateUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	isActiveStr := c.PostForm("is_active")

	// convert "true"/"false" string to bool
	isActive, _ := strconv.ParseBool(isActiveStr)

	// create user
	user := models.User{
		Username: username,
		Password: password, // ⚠️ should hash before saving!
		Email:    email,
		IsActive: isActive,
	}

	result := configs.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// redirect back to user list or show success
	c.Redirect(http.StatusSeeOther, "/admin/users")
}

func EditUserForm(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := configs.DB.First(&user, id).Error; err != nil {
		c.String(http.StatusNotFound, "User not found")
		return
	}

	c.HTML(http.StatusOK, "edituser.html", gin.H{
		"user": user,
	})
}

func UpdateUser(c *gin.Context) {
	// Load user first
	var user models.User
	if err := configs.DB.First(&user).Error; err != nil {
		c.String(http.StatusNotFound, "User not found")
		return
	}

	// Temporary struct to bind form values
	var form models.User
	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusBadRequest, "Error: %v", err)
		return
	}

	// Update only the fields from the form
	user.Username = form.Username
	user.Email = form.Email

	configs.DB.Save(&user)
	c.Redirect(http.StatusFound, "/admin/users")

}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := configs.DB.Delete(&models.User{}, id).Error; err != nil {
		c.String(http.StatusInternalServerError, "Failed to delete user")
		return
	}

	// After delete, go back to users list
	c.Redirect(http.StatusFound, "/admin/users")
}

// Update user
// func UpdateUser(c *gin.Context) {
// 	id := c.Param("id")
// 	var user models.User
// 	if err := configs.DB.First(&user, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}

// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	configs.DB.Save(&user)
// 	c.JSON(http.StatusOK, user)
// }

// Delete user
// func DeleteUser(c *gin.Context) {
// 	id := c.Param("id")
// 	if err := configs.DB.Delete(&models.User{}, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
// }

func CreateUsers(c *gin.Context) {
	c.HTML(http.StatusOK, "create.html", nil)
}
