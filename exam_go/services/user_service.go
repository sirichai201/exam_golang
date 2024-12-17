package services

import (
	"errors"          // ใช้สำหรับสร้างและจัดการ error
	"exam_go/models"   // นำเข้า model ของ User_test
)

// กำหนด error สำหรับกรณีที่ไม่พบผู้ใช้
var ErrUserNotFound = errors.New("user not found")

// ใช้ slice เก็บข้อมูลผู้ใช้ในหน่วยความจำ
var users []models.User_test

// CreateUser adds a new user to the list
// ฟังก์ชันสำหรับเพิ่มผู้ใช้ใหม่เข้าไปใน slice
func CreateUser(user models.User_test) (models.User_test, error) {
	users = append(users, user) // เพิ่มผู้ใช้ใหม่เข้าไปใน slice
	return user, nil            // คืนค่าผู้ใช้และ error เป็น nil
}

// GetUsers returns all users
// ฟังก์ชันสำหรับดึงข้อมูลผู้ใช้ทั้งหมด
func GetUsers() ([]models.User_test, error) {
	return users, nil // คืนค่า slice ของผู้ใช้ทั้งหมด และ error เป็น nil
}

// GetUserByID finds a user by ID
// ฟังก์ชันสำหรับค้นหาผู้ใช้ตาม ID
func GetUserByID(id uint) (models.User_test, error) {
	for _, user := range users { // วนลูปดูผู้ใช้ใน slice
		if user.ID == id { // ถ้าเจอผู้ใช้ที่มี ID ตรงกัน
			return user, nil // คืนค่าผู้ใช้นั้นและ error เป็น nil
		}
	}
	return models.User_test{}, ErrUserNotFound // กรณีไม่พบผู้ใช้ ให้คืนค่าข้อผิดพลาด
}

// UpdateUser updates an existing user's details
// ฟังก์ชันสำหรับอัปเดตข้อมูลผู้ใช้
func UpdateUser(id uint, updatedUser models.User_test) (models.User_test, error) {
	for index, user := range users { // วนลูปค้นหาผู้ใช้ใน slice
		if user.ID == id { // ถ้าเจอผู้ใช้ที่มี ID ตรงกัน
			users[index] = updatedUser // อัปเดตข้อมูลผู้ใช้ในตำแหน่งนั้น
			return updatedUser, nil    // คืนค่าผู้ใช้ที่ถูกอัปเดตและ error เป็น nil
		}
	}
	return models.User_test{}, ErrUserNotFound // กรณีไม่พบผู้ใช้ ให้คืนค่าข้อผิดพลาด
}

// DeleteUser deletes a user by ID
// ฟังก์ชันสำหรับลบผู้ใช้ตาม ID
func DeleteUser(id uint) error {
	for index, user := range users { // วนลูปค้นหาผู้ใช้ใน slice
		if user.ID == id { // ถ้าเจอผู้ใช้ที่มี ID ตรงกัน
			users = append(users[:index], users[index+1:]...) // ลบผู้ใช้ออกจาก slice
			return nil // คืนค่า nil เพื่อบอกว่าการลบสำเร็จ
		}
	}
	return ErrUserNotFound // กรณีไม่พบผู้ใช้ ให้คืนค่าข้อผิดพลาด
}
