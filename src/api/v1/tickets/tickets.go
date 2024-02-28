package tickets

import "github.com/gofiber/fiber/v2"

func TicketsRouter(router fiber.Router) {
	ticketsURL := router.Group("/tickets")

	ticketsURL.Post("")
	ticketsURL.Get("")
	ticketsURL.Get("/:id")
	ticketsURL.Put("/:id")
	ticketsURL.Delete("/:id")
}
