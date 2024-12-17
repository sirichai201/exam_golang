package controllers

import (
	"exam_go/models"
	us "exam_go/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ฟังก์ชันสำหรับการสมัครสมาชิก (Register)
func RegisterUser(c *gin.Context) {
    var user models.RegisterUsers

    // รับข้อมูลจาก JSON request body
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "Invalid input",
            "success": false,
            "code":    http.StatusBadRequest,
        })
        return
    }

    // เรียกใช้ service เพื่อทำการสมัครผู้ใช้
    if err := us.RegisterUser(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": err.Error(),
            "success": false,
            "code":    http.StatusBadRequest,
        })
        return
    }

    // ส่งข้อมูลที่มี account_id กลับไปใน response
    c.JSON(http.StatusOK, gin.H{
        "account_id": user.Acount_ID,
        "username":   user.Username,
        "name":       user.Name,
        "email":      user.Email,
        "password":   user.Password, 
        "success":    true,
        "code":       http.StatusOK,
    })
}

// Login: ฟังก์ชันสำหรับการล็อกอิน
func Login(c *gin.Context) {
    var credentials models.RegisterUsers

    // รับข้อมูลจาก JSON request body
    if err := c.ShouldBindJSON(&credentials); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "Invalid input",
            "success": false,
            "code":    http.StatusBadRequest,
        })
        return
    }

    // ตรวจสอบผู้ใช้
    user, err := us.AuthenticateUser(credentials)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "message": err.Error(),
            "success": false,
            "code":    http.StatusUnauthorized,
        })
        return
    }

    // สร้าง JWT token
    token, err := us.GenerateToken(user)
    if err != nil {
        log.Printf("Error generating token for user %s: %v\n", user.Username, err) // Log สำหรับ debug
        c.JSON(http.StatusInternalServerError, gin.H{
            "message": "Failed to generate token",
            "success": false,
            "code":    http.StatusInternalServerError,
        })
        return
    }

    // ส่ง Token กลับ
    c.JSON(http.StatusOK, gin.H{
        "token":   token,
        "success": true,
        "code":    http.StatusOK,
    })
}




func GetTestAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ทดสอบ Auth",
	})
}