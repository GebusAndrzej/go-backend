package v1

import (
	"go-backend/src/api/v1/products"
	"go-backend/src/api/v1/tickets"

	"github.com/gofiber/fiber/v2"
)

func ApiV1Router(router fiber.Router) {
	v1URL := router.Group("/v1")

	products.ProductsRouter(v1URL)
	tickets.TicketsRouter(v1URL)
}
