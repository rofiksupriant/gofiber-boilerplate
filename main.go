package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})

	app.Get("/:name", func(ctx *fiber.Ctx) error {
		name := ctx.Params("name")
		return ctx.SendString("Hello " + name)
	})

	_ = app.Listen(":3000")
}
