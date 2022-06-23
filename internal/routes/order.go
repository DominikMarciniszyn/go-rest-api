package routes

import (
	"go-rest-api/internal/database/config"
	"go-rest-api/internal/database/entities"

	"github.com/gofiber/fiber"
)

func GetOrders(context *fiber.Ctx) {
	var orders []entities.Order

	config.Database.Find(&orders)

	context.Status(200).JSON(&orders)
}

func GetOrder(context *fiber.Ctx) {
	id := context.Params("id")
	var order entities.Order

	result := config.Database.Find(&order, id)

	if result.RowsAffected == 0 {
		context.Status(404)
	}

	context.Status(200).JSON(&order)
}

func CreateOrder(context *fiber.Ctx) {
	order := new(entities.Order)

	if context.BodyParser(order) != nil {
		context.Status(503).Error()
	}

	config.Database.Create(&order)
	context.Status(201).JSON(&order)
}

func UpdateOrder(context *fiber.Ctx) {
	id := context.Params("id")
	order := new(entities.Order)

	if context.BodyParser(order) != nil {
		context.Status(503).Error()
	}

	config.Database.Where("id = ?", id).Updates(&order)
	context.Status(200).JSON(order)
}

func RemoveOrder(context *fiber.Ctx) {
	id := context.Params("id")
	var order entities.Order

	result := config.Database.Delete(&order, id)

	if result.RowsAffected == 0 {
		context.Status(404).Error()
	}

	context.Status(200).JSON("Deleted")
}
