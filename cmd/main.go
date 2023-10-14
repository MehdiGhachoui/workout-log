package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mehdighachoui/workout-log/config"
)

// type Log struct {
// 	Weight    float64
// 	Type      string
// 	Exercices []*Exercice
// }

// {number}gg = jump to line number
// L : move to next page ?
// H : move to prev page ?

// type Exercice struct {
// 	Name   string
// 	Goal   string
// 	Reps   []int64
// 	Weight []float64
// }

func main() {

	config.Connect()
	log.Println("Database connected")

	app := fiber.New()
	app.Use(logger.New())
	app.Get("/healthCheck", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	log.Fatal(app.Listen(":8000"))

}
