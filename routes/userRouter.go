package routes

import (
	"restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:id", controllers.GetUser())
	incomingRoutes.POST("/users/signup", controllers.CreateUser())
	incomingRoutes.POST("/users/login", controllers.LoginUser())
}
