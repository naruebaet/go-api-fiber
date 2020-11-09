package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(&fiber.Map{
			"s":   "ok",
			"msg": "hello world",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
