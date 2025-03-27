package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOrdersItems() fiber.Handler {
	return func(c fiber.Ctx) error {

	}
}

func ItemsByOrder(id string) (OrderItems []primitive.M, err error) {

}
func GetOrderItem() fiber.Handler {
	return func(c fiber.Ctx) error {

	}
}

func GetOrderItemsByOrder() fiber.Handler {
	return func(c fiber.Ctx) error {

	}
}

func CreateOrderItem() fiber.Handler {
	return func(c fiber.Ctx) error {

	}
}

func UpdateOrderItem() fiber.Handler {
	return func(c fiber.Ctx) error {

	}
}
