package routes

import (
	"restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", controllers.GetOrders())
	incomingRoutes.GET("/orders/:id", controllers.GetOrder())
	incomingRoutes.POST("/orders", controllers.CreateOrder())
	incomingRoutes.PATCH("/orders/:id", controllers.UpdateOrder())
	incomingRoutes.DELETE("/orders/:id", controllers.DeleteOrder())
}
