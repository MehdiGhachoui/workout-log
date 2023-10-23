package pkg

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func WorkoutHTTP(f *fiber.App) {

	f.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})
	f.Post("/", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})
	f.Put("/", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})
	f.Delete("/", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})
}
