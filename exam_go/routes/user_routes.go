package routes

import (
	uc "exam_go/controllers"
	"exam_go/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/signup", uc.RegisterUser)
		user.POST("/login", uc.Login)

	}

	// Protected routes
	protected := r.Group("/protected", middleware.AuthRequired())
	{
		protected.GET("/GetTestAuth", uc.GetTestAuth)
	}
}
