package controllers

import (
	"establishment/internal/models"
	usecasescategory "establishment/internal/usecases/category"

	"github.com/gofiber/fiber/v2"
)

func CreateCategory(c *fiber.Ctx) error {
	var category models.Category

	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Requisição inválida"})
	}

	err := usecasescategory.CreateCategory(category)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao criar estabelecimento"})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Estabelecimento criado com sucesso!"})
}
