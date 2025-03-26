package routes

import (
	controller "Resturant-Management-System/controllers"
	"github.com/gofiber/fiber/v2"
)

func InvoiceRoutes(incomingRoutes *fiber.App) {
	incomingRoutes.GET("/invoices", controller.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", controller.GetInvoice())
	incomingRoutes.POST("/invoices", controller.CreateInvoice())
	incomingRoutes.PATCH("/invoices/:invoice_id", controller.UpdateInvoice())
}
