package routes

import (
	"github.com/gin-gonic/gin"
	uc "exam_go/controllers"
)

func UserRoutes(r *gin.Engine) { 
	usergroups := r.Group("/users")

	usergroups.POST("", uc.CreateUser)
	usergroups.GET("", uc.GetUsers)
	usergroups.GET("/:id", uc.GetUser)
	usergroups.PUT("/:id", uc.UpdateUser)
	usergroups.DELETE("/:id", uc.DeleteUser)
}
