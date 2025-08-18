package databases

import (
	"Day10/configs"
	"Day10/models"
	"fmt"
)

func AutoMigrate() {
	err := configs.DB.AutoMigrate(
		&models.Product{},
		&models.User{},
	)
	if err != nil {
		errorLog := fmt.Sprintf("Gagal Auto Migrate: %s", err.Error())
		panic(errorLog)
	}
}
