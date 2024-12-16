package services

import (
	"exam_go/models"
	"net/http"
)

var users []models.User_test

// CreateUser adds a new user to the list
func CreateUser(user models.User_test) (models.User_test, error) {
	users = append(users, user)
	return user, nil
}

// GetUsers returns all users
func GetUsers() ([]models.User_test, error) {
	return users, nil
}

// GetUserByID finds a user by ID
func GetUserByID(id uint) (models.User_test, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return models.User_test{}, http.ErrNotFound
}

// UpdateUser updates an existing user's details
func UpdateUser(id uint, updatedUser models.User_test) (models.User_test, error) {
	for index, user := range users {
		if user.ID == id {
			users[index] = updatedUser
			return updatedUser, nil
		}
	}
	return models.User_test{}, http.ErrNotFound 
}

// DeleteUser deletes a user by ID
func DeleteUser(id uint) error {
	for index, user := range users {
		if user.ID == id {
			users = append(users[:index], users[index+1:]...)
			return nil
		}
	}
	return http.ErrNotFound
}
