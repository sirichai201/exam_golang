package magrate

import (
	"exam_go/config"
	"exam_go/models"
	"log"
)

// AutoMigrateTables automates the migration of database tables
func AutoMigrateTables() {
	// เพิ่มตารางที่ต้องการ AutoMigrate
	tables := []interface{}{
		&models.User_test{},
	}

	// AutoMigrate สำหรับทุกตาราง
	for _, table := range tables {
		err := config.DB.AutoMigrate(table)
		if err != nil {
			log.Fatalf("Failed to migrate table %T: %v", table, err)
		}
	}
	log.Println("Database migration completed successfully!")
}
