package models

import "gorm.io/gorm"

type RegisterUsers struct {
	gorm.Model        // ใช้ gorm.Model สำหรับการจัดการเวลา CreatedAt, UpdatedAt, DeletedAt
	Acount_ID  uint   `json:"acount_id"`
	Username   string `json:"username" gorm:"unique"` // ระบุว่า Username ต้องเป็น unique
	Name       string `json:"name"`
	Email      string `json:"email" gorm:"unique"` // ระบุว่า Email ต้องเป็น unique
	Password   string `json:"-"`                   // ไม่ให้แสดงรหัสผ่านใน JSON
}
