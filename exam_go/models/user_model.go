package models

import "gorm.io/gorm"

type User_test struct {
	gorm.Model
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
}
