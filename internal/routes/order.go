package routes

import (
	"github.com/gofiber/fiber"
)

func GetOrders(context *fiber.Ctx) {
	context.Status(200).SendString("All orders")
}

func GetOrder(context *fiber.Ctx) {

}

func CreateOrder(context *fiber.Ctx) {

}
