package routes

import (
	"exam_go/controllers"
	"exam_go/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/signup", controllers.RegisterUser)
		user.POST("/login", controllers.Login)

	}

	protected := r.Group("/protected", middleware.AuthRequired()) // ใช้ AuthRequired Middleware
    {
        protected.GET("/GetTestAuth", controllers.GetTestAuth)
    }
}


