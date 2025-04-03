// internal/controllers/user_controller.go
package controllers

import (
	"establishment/internal/models"
	"establishment/internal/usecases"
	"fmt"

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

func ListEstablishments(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	users, total, err := usecases.GetEstablishmentWithPagination(page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao buscar usuários",
		})
	}

	lastPage := (total + limit - 1) / limit

	return c.JSON(fiber.Map{
		"data": users,
		"meta": fiber.Map{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"totalPages": lastPage,
		},
		"links": fiber.Map{
			"self": fmt.Sprintf("/api/users?page=%d&limit=%d", page, limit),
			"next": fmt.Sprintf("/api/users?page=%d&limit=%d", page+1, limit),
			"prev": fmt.Sprintf("/api/users?page=%d&limit=%d", max(page-1, 1), limit),
		},
	})
}

// Util para evitar página 0
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
