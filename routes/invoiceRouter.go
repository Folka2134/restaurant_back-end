package routes

import (
	"restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controllers.GetInvoices())
	incomingRoutes.GET("/invoices/:id", controllers.GetInvoice())
	incomingRoutes.POST("/invoices", controllers.CreateInvoice())
	incomingRoutes.PATCH("/invoices/:id", controllers.UpdateInvoice())
	incomingRoutes.DELETE("/invoices/:id", controllers.DeleteInvoice())
}
