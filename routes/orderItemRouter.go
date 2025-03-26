package routes

import (
	controller "Resturant-Management-System/controllers"
	"github.com/gofiber/fiber/v2"
)

func OrderItemRoutes(incomingRoutes *fiber.App) {
	incomingRoutes.GET("/ordersItems", controller.GetOrdersItems())
	incomingRoutes.GET("/ordersItems/:orderItem_id", controller.GetOrderItem())
	incomingRoutes.GET("orderItems-order/:order_id", controller.GetOrderItemsByOrder())
	incomingRoutes.POST("/ordersItems", controller.CreateOrderItem())
	incomingRoutes.PATCH("/ordersItems/:orderItem_id", controller.UpdateOrderItem())
}
