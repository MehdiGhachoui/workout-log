package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mehdighachoui/workout-log/db"
	"github.com/mehdighachoui/workout-log/pkg"
)

func main() {
	_, err := db.CreatePostgresConnection()

	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	log.Println("Database connected")

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/healthCheck", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	pkg.WorkoutHTTP(app)
	log.Fatal(app.Listen(":8000"))
}
