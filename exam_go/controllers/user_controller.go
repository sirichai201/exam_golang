package controllers


import (
    "github.com/gin-gonic/gin"
    "exam_go/models"
    us "exam_go/services"
    "net/http"
)

// CreateUser handles creating a new user
func CreateUser(c *gin.Context) {
    var newUser models.User_test
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := us.CreateUser(newUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, user)
}

// GetUsers handles retrieving all users
func GetUsers(c *gin.Context) {
    users, err := us.GetUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}

// GetUser handles retrieving a single user by ID
func GetUser(c *gin.Context) {
    id := c.Param("id")
    user, err := us.GetUserByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

// UpdateUser handles updating an existing user
func UpdateUser(c *gin.Context) {
    id := c.Param("id")
    var updatedUser models.User_test
    if err := c.ShouldBindJSON(&updatedUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := us.UpdateUser(id, updatedUser)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

// DeleteUser handles deleting a user by ID
func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    err := us.DeleteUser(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
