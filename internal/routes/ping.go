package routes

import "github.com/gofiber/fiber"

func Ping(context *fiber.Ctx) {
	context.Status(200).SendString("Service is alive...")
}
