package routes

import (
	controller "Resturant-Management-System/controllers"
	"github.com/gofiber/fiber/v2"
)

func TableRoutes(incomingRoutes *fiber.App) {
	incomingRoutes.GET("/tables", controller.GetTables())
	incomingRoutes.GET("/tables/:table_id", controller.GetTable())
	incomingRoutes.POST("/tables", controller.CreateTable())
	incomingRoutes.PATCH("/tables/:tables_id", controller.UpdateTable())
}
