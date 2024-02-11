package routes

import (
	"restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menu", controllers.GetMenus())
	incomingRoutes.GET("/menu/:id", controllers.GetMenu())
	incomingRoutes.POST("/menu", controllers.CreateMenu())
	incomingRoutes.PATCH("/menu/:id", controllers.UpdateMenu())
	incomingRoutes.DELETE("/menu/:id", controllers.DeleteMenu())
}
