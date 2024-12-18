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
		tokenString := c.GetHeader("Authorization")

		// ตรวจสอบ token ว่ามีคำว่า Bearer หรือไม่ และตัดออก
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"success": false,
				"code":    http.StatusUnauthorized,
			})
			c.Abort()
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// ตรวจสอบและแยกข้อมูลใน token
		claims, err := services.ValidateAndExtractToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
				"success": false,
				"code":    http.StatusUnauthorized,
			})
			c.Abort()
			return 
		}

		c.Set("id", claims["id"])
		c.Set("account_id", claims["account_id"])
        c.Set("user_email", claims["email"])
        c.Set("user_username", claims["username"])  
        c.Set("user_name", claims["name"])
		c.Set("phone", claims["phone"])

		// ผ่าน middleware
		c.Next()
	}
}




