// internal/routes/routes.go
package routes

import (
	"establishment/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/establishment", controllers.CreateEstablishment)
	api.Get("/establishment", controllers.ListEstablishments)
}
