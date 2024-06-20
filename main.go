package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ssoda/fibercaptcha"
)

func main() {
	app := fiber.New()
	app.Use(fibercaptcha.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
