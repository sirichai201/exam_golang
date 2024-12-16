package main

import (
	"exam_go/database"
	"exam_go/models"
  	ur "exam_go/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// เชื่อมต่อฐานข้อมูล
	database.ConnectDB()

	// ทำ Auto Migrate
	database.DB.AutoMigrate(&models.User_test{})

	// สร้าง Gin Router
	r := gin.Default()

	// ลงทะเบียน Routes
	ur.UserRoutes(r)

	// เริ่มเซิร์ฟเวอร์
	log.Println("Server is running on http://localhost:8080")
	r.Run(":8080")
}
