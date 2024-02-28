package products

import "github.com/gofiber/fiber/v2"

func ProductsRouter(router fiber.Router) {
	productsURL := router.Group("/products")

	productsURL.Post("")
	productsURL.Get("")
	productsURL.Get("/:id")
	productsURL.Put("/:id")
	productsURL.Delete("/:id")
}
