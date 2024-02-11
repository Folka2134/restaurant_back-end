package routes

import (
	"restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controllers.GetOrderItems())
	incomingRoutes.GET("/orderItems/:id", controllers.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:order_id", controllers.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItem", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:id", controllers.UpdateOrderItem())
	incomingRoutes.DELETE("/orderItems/:id", controllers.DeleteOrderItem())
}
