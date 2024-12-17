package controllers

import (
    "exam_go/models"
    us "exam_go/services"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

// CreateUser handles creating a new user
// ฟังก์ชันสำหรับสร้างผู้ใช้ใหม่
func CreateUser(c *gin.Context) {
    var newUser models.User_test // ตัวแปรสำหรับเก็บข้อมูลผู้ใช้ใหม่จาก JSON
    if err := c.ShouldBindJSON(&newUser); err != nil { // ตรวจสอบและแปลง JSON ที่ส่งมาจากผู้ใช้เป็น Struct
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // ส่งข้อความผิดพลาดกรณี JSON ไม่ถูกต้อง
        return
    }

    user, err := us.CreateUser(newUser) // เรียกใช้บริการเพื่อสร้างผู้ใช้
    if err != nil { // ตรวจสอบว่ามีข้อผิดพลาดระหว่างการสร้างหรือไม่
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // ส่งข้อความผิดพลาดกรณีสร้างไม่สำเร็จ
        return
    }
    c.JSON(http.StatusOK, user) // ส่งข้อมูลผู้ใช้ที่สร้างสำเร็จกลับไป
}


// GetUsers handles retrieving all users
// ฟังก์ชันสำหรับดึงข้อมูลผู้ใช้ทั้งหมด
func GetUsers(c *gin.Context) {
    users, err := us.GetUsers() // เรียกใช้บริการเพื่อดึงข้อมูลผู้ใช้ทั้งหมด
    if err != nil { // ตรวจสอบว่ามีข้อผิดพลาดหรือไม่
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // ส่งข้อความผิดพลาดกรณีเกิดปัญหา
        return
    }
    c.JSON(http.StatusOK, users) // ส่งข้อมูลผู้ใช้ทั้งหมดกลับไป
}

// GetUser handles retrieving a single user by ID
// ฟังก์ชันสำหรับดึงข้อมูลผู้ใช้ตาม ID
func GetUser(c *gin.Context) {
    idStr := c.Param("id") // ดึงค่า ID จาก URL Parameter
    id, err := strconv.ParseUint(idStr, 10, 32) // แปลง ID จาก string เป็น uint
    if err != nil { // ตรวจสอบว่าการแปลงสำเร็จหรือไม่
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"}) // ส่งข้อความผิดพลาดกรณี ID ไม่ถูกต้อง
        return
    }

    user, err := us.GetUserByID(uint(id)) // เรียกใช้บริการเพื่อดึงข้อมูลผู้ใช้ตาม ID
    if err != nil { // ตรวจสอบว่าพบผู้ใช้หรือไม่
        c.JSON(http.StatusNotFound, gin.H{"message": "User not found"}) // ส่งข้อความว่าผู้ใช้ไม่พบ
        return
    }
    c.JSON(http.StatusOK, user) // ส่งข้อมูลผู้ใช้กลับไป
}

// UpdateUser handles updating an existing user
// ฟังก์ชันสำหรับแก้ไขข้อมูลผู้ใช้
func UpdateUser(c *gin.Context) {
    idStr := c.Param("id") // ดึงค่า ID จาก URL Parameter
    id, err := strconv.ParseUint(idStr, 10, 32) // แปลง ID จาก string เป็น uint
    if err != nil { // ตรวจสอบว่าการแปลงสำเร็จหรือไม่
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"}) // ส่งข้อความผิดพลาดกรณี ID ไม่ถูกต้อง
        return
    }

    var updatedUser models.User_test // ตัวแปรสำหรับเก็บข้อมูลผู้ใช้ใหม่จาก JSON
    if err := c.ShouldBindJSON(&updatedUser); err != nil { // ตรวจสอบและแปลง JSON ที่ส่งมาจากผู้ใช้เป็น Struct
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // ส่งข้อความผิดพลาดกรณี JSON ไม่ถูกต้อง
        return
    }

    user, err := us.UpdateUser(uint(id), updatedUser) // เรียกใช้บริการเพื่อแก้ไขข้อมูลผู้ใช้
    if err != nil { // ตรวจสอบว่าพบผู้ใช้หรือไม่
        c.JSON(http.StatusNotFound, gin.H{"message": "User not found"}) // ส่งข้อความว่าผู้ใช้ไม่พบ
        return
    }
    c.JSON(http.StatusOK, user) // ส่งข้อมูลผู้ใช้ที่ถูกแก้ไขกลับไป
}

// DeleteUser handles deleting a user by ID
// ฟังก์ชันสำหรับลบผู้ใช้ตาม ID
func DeleteUser(c *gin.Context) {
    idStr := c.Param("id") // ดึงค่า ID จาก URL Parameter
    id, err := strconv.ParseUint(idStr, 10, 32) // แปลง ID จาก string เป็น uint
    if err != nil { // ตรวจสอบว่าการแปลงสำเร็จหรือไม่
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"}) // ส่งข้อความผิดพลาดกรณี ID ไม่ถูกต้อง
        return
    }

    err = us.DeleteUser(uint(id)) // เรียกใช้บริการเพื่อลบผู้ใช้
    if err != nil { // ตรวจสอบว่าพบผู้ใช้หรือไม่
        c.JSON(http.StatusNotFound, gin.H{"message": "User not found"}) // ส่งข้อความว่าผู้ใช้ไม่พบ
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User deleted"}) // ส่งข้อความยืนยันการลบผู้ใช้สำเร็จ
}
