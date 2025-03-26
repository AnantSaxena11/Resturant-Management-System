package routes

import (
	controller "Resturant-Management-System/controllers"
	"github.com/gofiber/fiber/v2"
)

func OrderRoutes(incomingRoutes *fiber.App) {
	incomingRoutes.GET("/orders", controller.GetOrders())
	incomingRoutes.GET("/orders/:order_id", controller.GetOrder())
	incomingRoutes.POST("/orders", controller.CreateOrder())
	incomingRoutes.PATCH("/orders/:order_id", controller.UpdateOrder())
}
