package main

import (
	"Resturant-Management-System/database"
	"Resturant-Management-System/middleware"
	"Resturant-Management-System/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	// # HERE WE ARE GETTING THE VARIABLES FROM THE ENV FILE
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found using default values")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
	routes.UserRoutes(app)
	app.Use(middleware.Authentication())
	routes.FoodRoutes(app)
	routes.MenuRoutes(app)
	routes.TableRoutes(app)
	routes.OrderRoutes(app)
	routes.OrderItemRoutes(app)
	routes.InvoiceRoutes(app)
	err = app.Listen(":" + port)
	if err != nil {
		log.Fatal(err)
	}

}
