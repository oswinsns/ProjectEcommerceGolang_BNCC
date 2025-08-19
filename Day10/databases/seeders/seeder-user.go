package seeders

import (
	"Day10/configs"
	"Day10/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// helper to hash password
func hashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}

func SeedUsers() {
	users := []models.User{
		{Username: "admin", Password: hashPassword("admin123"), Email: "admin@example.com", IsActive: true},
		{Username: "johndoe", Password: hashPassword("password123"), Email: "john@example.com", IsActive: true},
		{Username: "janedoe", Password: hashPassword("mypassword"), Email: "jane@example.com", IsActive: false},
	}

	// 🔥 Fully clear table first
	configs.DB.Exec("TRUNCATE TABLE users")

	// Insert fresh seed data
	if err := configs.DB.Create(&users).Error; err != nil {
		log.Fatalf("failed seeding users: %v", err)
	}

	log.Println("Re-seeded users successfully 🚀")

	// var count int64
	// configs.DB.Model(&models.User{}).Count(&count)
	// if count == 0 {
	// 	if err := configs.DB.Create(&users).Error; err != nil {
	// 		log.Fatalf("failed seeding users: %v", err)
	// 	}
	// 	log.Println("Seeded users successfully 🚀")
	// } else {
	// 	log.Println("Users already exist, skipping seeding.")
	// }

}
