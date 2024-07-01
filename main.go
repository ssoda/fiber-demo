package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/ssoda/fibercaptcha"
)

type LoginInput struct {
	CaptchaID string `json:"captcha_id"`
	Captcha   string `json:"captcha"`
}

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(fibercaptcha.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		input := new(LoginInput)
		if err := c.BodyParser(input); err != nil {
			log.Println("parse input error:", err)
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.SendString("ok")
	})

	log.Fatal(app.Listen(":3000"))
}
