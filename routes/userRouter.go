package routes

import (
	controller "Resturant-Management-System/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(incomingRoutes *fiber.App) {
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/user_id", controller.GetUser())
	incomingRoutes.POST("/users/signup", controller.SignUp())
	incomingRoutes.POST("/users/signin", controller.SignIn())
}
