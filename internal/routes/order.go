package routes

import (
	"go-rest-api/internal/database/config"
	"go-rest-api/internal/database/entities"

	"github.com/gofiber/fiber"
)

func GetOrders(context *fiber.Ctx) error {
	var orders []entities.Order

	config.Database.Find(&orders)

	return context.Status(200).JSON(&orders)
}

func GetOrder(context *fiber.Ctx) error {
	id := context.Params("id")
	var order entities.Order

	result := config.Database.Find(&order, id)

	if result.RowsAffected == 0 {
		return context.Status(404).Error()
	}

	return context.Status(200).JSON(&order)
}

func CreateOrder(context *fiber.Ctx) error {
	order := new(entities.Order)

	if context.BodyParser(order) != nil {
		return context.Status(503).Error()
	}

	config.Database.Create(&order)
	return context.Status(201).JSON(&order)
}

func UpdateOrder(context *fiber.Ctx) error {
	id := context.Params("id")
	order := new(entities.Order)

	if context.BodyParser(order) != nil {
		return context.Status(503).Error()
	}

	config.Database.Where("id = ?", id).Updates(&order)
	return context.Status(200).JSON(order)
}

func RemoveOrder(context *fiber.Ctx) error {
	id := context.Params("id")
	var order entities.Order

	result := config.Database.Delete(&order, id)

	if result.RowsAffected == 0 {
		return context.Status(404).Error()
	}

	return context.Status(200).JSON("Deleted")
}
