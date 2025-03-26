package routes

import (
	controller "Resturant-Management-System/controllers"
	"github.com/gofiber/fiber/v2"
)

func FoodRoutes(incomingRoutes *fiber.App) {
	incomingRoutes.GET("/food", controller.GetFoods())
	incomingRoutes.GET("/food/:food_id", controller.GetFood())
	incomingRoutes.POST("/foods", controller.CreateFood())
	incomingRoutes.PATCH("/foods/:food_id", controller.UpdateFood())
}
