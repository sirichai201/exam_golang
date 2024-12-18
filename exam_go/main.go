package main

import (
	"exam_go/config"
  	ur "exam_go/routes"
	"log"
	"exam_go/magrate"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	// เชื่อมต่อฐานข้อมูล
	config.ConnectDB()

	// ทำ Auto Migrate
	magrate.AutoMigrateTables()

	// สร้าง Gin Router
	r := gin.Default()

	// ลงทะเบียน Routes
	ur.UserRoutes(r)

	// เริ่มเซิร์ฟเวอร์
	log.Println("Server is running on http://localhost:8080")
	r.Run(":8080")
}
