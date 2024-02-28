package auth

import "github.com/gofiber/fiber/v2"

func AuthRouter(router fiber.Router) {
	productsURL := router.Group("/auth")

	productsURL.Post("/login", loginMock)
}

func loginMock(c *fiber.Ctx) error {
	return c.Status(200).SendString("logged")
}
