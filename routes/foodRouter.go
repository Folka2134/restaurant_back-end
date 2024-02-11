package routes

import (
	"restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/food", controllers.GetFoods())
	incomingRoutes.GET("/food/:id", controllers.GetFood())
	incomingRoutes.POST("/food", controllers.CreateFood())
	incomingRoutes.PATCH("/food/:id", controllers.UpdateFood())
	incomingRoutes.DELETE("/food/:id", controllers.DeleteFood())
}
