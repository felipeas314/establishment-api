// internal/controllers/user_controller.go
package controllers

import (
	"establishment/internal/models"
	"establishment/internal/usecases"

	"github.com/gofiber/fiber/v2"
)

// Criar um usuário
func CreateEstablishment(c *fiber.Ctx) error {
	var establishment models.Establishment

	if err := c.BodyParser(&establishment); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Requisição inválida"})
	}

	err := usecases.CreateEstablishment(establishment)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao criar estabelecimento"})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Estabelecimento criado com sucesso!"})
}
