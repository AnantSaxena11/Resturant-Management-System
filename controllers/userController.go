package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetUsers() fiber.Handler {
	return func(c fiber.Ctx) error {

	}
}

func GetUser() fiber.Handler {
	return func(c fiber.Ctx) error {

	}
}

func SignUp() fiber.Handler {
	return func(c fiber.Ctx) error {

	}
}

func SignIn() fiber.Handler {
	return func(c fiber.Ctx) error {

	}
}

func HashedPassword(password string) string {

}

func VerifyPassword(userPassword string, providePassword string) (bool, string) {

}
