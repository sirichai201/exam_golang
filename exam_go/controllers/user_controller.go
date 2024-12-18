package controllers

import (
	"exam_go/models"
	 "exam_go/services"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "fmt"
)

// ฟังก์ชันสำหรับการสมัครสมาชิก (Register)
func RegisterUser(c *gin.Context) {
    var user models.Users

    // รับข้อมูลจาก JSON request body
    if err := c.ShouldBindJSON(&user); err != nil {
        // ตรวจสอบ validation errors
        validationErrors, ok := err.(validator.ValidationErrors)
        if ok {
            errors := make(map[string]string)
            for _, err := range validationErrors {
                field := err.Field()
                tag := err.Tag()
                errors[field] = fmt.Sprintf("must be %s", tag)
            }

            c.JSON(http.StatusBadRequest, gin.H{
                "message": "Validation failed",
                "errors":  errors,
                "success": false,
                "code":    http.StatusBadRequest,
            })
            return
        }

        // กรณีมีข้อผิดพลาดทั่วไป
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "Invalid input",
            "success": false,
            "code":    http.StatusBadRequest,
        })
        return
    }

    // เรียกใช้ service เพื่อทำการสมัครผู้ใช้
    if err := services.RegisterUser(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": err.Error(),
            "success": false,
            "code":    http.StatusBadRequest,
        })
        return
    }

    // ส่งข้อมูลที่มี account_id กลับไปใน response
    c.JSON(http.StatusOK, gin.H{
        "message": "User registered successfully",
        "success": true,
        "code": http.StatusOK,
        "data": gin.H{
            "account_id":      user.AccountID,
            "address":         user.Address,
            "company_id":      user.CompanyID,
            "created_at":      user.CreatedAt,
            "created_by":      user.CreatedBy,
            "created_date":    user.CreatedDate,
            "deleted_at":      user.DeletedAt.Time,
            "email":           user.Email,
            "facebook":        user.Facebook,
            "mobile_number":   user.MobileNumber,
            "name":            user.Name,
            "one_id":          user.OneID,
            "prename":         user.Prename,
            "profile_picture": user.ProfilePicture,
            "surname":         user.Surname,
            "updated_at":      user.UpdatedAt,
            "updated_by":      user.UpdatedBy,
            "updated_date":    user.UpdatedDate,
            "username":        user.Username,
        },
    })
}

// Login: ฟังก์ชันสำหรับการล็อกอิน
func Login(c *gin.Context) {
    // กำหนดฟิลด์ที่ต้องการ
    requiredFields := []string{"username", "password"}

    // รับข้อมูลเป็น map[string]interface{}
    var input map[string]interface{}
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "Invalid JSON input",
            "success": false,
            "code":    http.StatusBadRequest,
        })
        return
    }

    // ตรวจสอบว่าฟิลด์ที่ต้องการมีอยู่หรือไม่
    missingFields := []string{}
    for _, field := range requiredFields {
        if _, exists := input[field]; !exists {
            missingFields = append(missingFields, field)
        }
    }

    // หากมีฟิลด์ที่ขาดไป
    if len(missingFields) > 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "Missing required fields",
            "missing_fields": missingFields,
            "success": false,
            "code":    http.StatusBadRequest,
        })
        return
    }

    username := input["username"].(string)
    password := input["password"].(string)

    

    // สมมุติเรียกใช้ service เพื่อยืนยันผู้ใช้
    user, err := services.AuthenticateUser(username, password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "message": err.Error(),
            "success": false,
            "code":    http.StatusUnauthorized,
        })
        return
    }

    // สร้าง JWT token
    token, err := services.GenerateToken(user)
    if err != nil {
        log.Printf("Error generating token for user %s: %v\n", username, err)
        c.JSON(http.StatusInternalServerError, gin.H{
            "message": "Failed to generate token",
            "success": false,
            "code":    http.StatusInternalServerError,
        })
        return
    }

    // ส่ง token กลับ
    c.JSON(http.StatusOK, gin.H{
        "data": gin.H{
            "token": token,
        },
        "success": true,
        "code":    http.StatusOK,
    })
}



func GetTestAuth(c *gin.Context) {
    // ดึงข้อมูลจาก context
    userAccount_id ,_ := c.Get("account_id")
    userEmail, _ := c.Get("user_email")
    userUsername, _ := c.Get("user_username")
    userName, _ := c.Get("user_name")
    userphone, _ := c.Get("phone")

    // จัดข้อมูลในรูปแบบ payload
    payload := gin.H{
        "userAccount_id":  userAccount_id,
        "email":    userEmail,
        "username": userUsername,
        "name":     userName,
        "phone":  userphone,
    }

    // ส่งข้อมูล payload กลับใน response
    c.JSON(http.StatusOK, gin.H{
        "message":  "ทดสอบ Auth",
        "success":  true,
        "code":     http.StatusOK,
        "data":     payload,  // ส่ง payload กลับไปใน response
    })
}



