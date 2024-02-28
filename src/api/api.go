package api

import (
	"go-backend/src/api/auth"

	"github.com/gofiber/fiber/v2"
)

func SetupApi(app *fiber.App) {
	app.Get("", healthCheck)
	api := app.Group("/api")

	auth.AuthRouter(api)
}

func healthCheck(c *fiber.Ctx) error {
	return c.Status(200).SendString("OK")
}
