// main.go
package main

import (
	"establishment/config"
	"establishment/internal/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()

	config.ConnectDB()

	app := fiber.New()

	routes.SetupRoutes(app)

	// ✅ 5. Inicia o servidor
	port := config.GetEnv("PORT", "8080")
	log.Fatal(app.Listen(":" + port))
}
