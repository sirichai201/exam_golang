package middleware

import (
	"exam_go/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ดึง token จาก Authorization header
		token := c.GetHeader("Authorization")

		// ลบ "Bearer " ออกจาก token
		token = strings.TrimPrefix(token, "Bearer ")

		// ตรวจสอบว่า token ว่างหรือ validate ไม่ผ่าน
		if token == "" || !services.ValidateToken(token) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"success": false,
				"code":    http.StatusUnauthorized,
			})
			c.Abort() // หยุดการดำเนินการ
			return
		}
		c.Next()
	}
}
