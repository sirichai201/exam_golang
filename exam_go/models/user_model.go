package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID             uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	AccountID      string    `json:"account_id" gorm:"type:varchar(255);unique"`
	OneID          string    `json:"one_id" gorm:"type:varchar(255)"`
	Username       string    `json:"username" gorm:"type:varchar(255);unique;not null" binding:"required"`
	Password       string    `json:"password" gorm:"type:longtext;not null" binding:"required"`
	CompanyID      string    `json:"company_id" gorm:"type:varchar(255);default:NULL"`
	Prename        string    `json:"prename" gorm:"type:varchar(255)"`
	Name           string    `json:"name" gorm:"type:varchar(255)"`
	Surname        string    `json:"surname" gorm:"type:varchar(255)"`
	Email          string    `json:"email" gorm:"type:varchar(255)" binding:"required,email"`
	MobileNumber   string    `json:"mobile_number" gorm:"type:varchar(20)" binding:"required"`
	Address        string    `json:"address" gorm:"type:text"`
	Facebook       string    `json:"facebook" gorm:"type:varchar(255)"`
	ProfilePicture string    `json:"profile_picture" gorm:"type:longtext"`
	CreatedDate    time.Time `json:"created_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	CreatedBy      string    `json:"created_by" gorm:"type:varchar(255);default:'EXAM_API_GO_REGISTER_USER'"`
	UpdatedDate    time.Time `json:"updated_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	UpdatedBy      string    `json:"updated_by" gorm:"type:varchar(255);default:NULL"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
