package magrate

import (
	"exam_go/config"
	"exam_go/models"
	"log"
)

func AutoMigrateTables() {
	err := config.DB.AutoMigrate(&models.Users{})
	if err != nil {
		log.Fatal("Failed to migrate tables:", err)
	}
	log.Println("Tables migrated successfully")
}