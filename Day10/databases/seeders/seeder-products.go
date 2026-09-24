package seeders

import (
	"Day10/configs"
	"Day10/models"
	"log"
)

func SeedProducts() {
	products := []models.Product{
		{Name: "Laptop", Stock: 10, Price: 15000000},
		{Name: "Mouse", Stock: 50, Price: 150000},
		{Name: "Keyboard", Stock: 30, Price: 450000},
		{Name: "Monitor", Stock: 15, Price: 2500000},
		{Name: "Headset", Stock: 20, Price: 750000},
	}

	// ❌ SEBELUMNYA (BUG): TRUNCATE menghapus 100% isi tabel produk setiap kali server restart
	// configs.DB.Exec("TRUNCATE TABLE products")
	// if err := configs.DB.Create(&products).Error; err != nil {
	// 	log.Fatalf("failed seeding products: %v", err)
	// }
	// log.Println("Re-seeded products successfully 🚀")

	// ✅ PERBAIKAN: Hanya isi data seeder jika tabel produk masih kosong (Count == 0)
	var count int64
	configs.DB.Model(&models.Product{}).Count(&count)
	if count == 0 {
		if err := configs.DB.Create(&products).Error; err != nil {
			log.Fatalf("failed seeding products: %v", err)
		}
		log.Println("Seeded products successfully 🚀")
	} else {
		log.Println("Products already exist, skipping seeding.")
	}
}
