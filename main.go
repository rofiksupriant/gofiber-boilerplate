package main

import (
	"boilerplate/internal/config"
	"boilerplate/internal/http"
	"boilerplate/internal/middleware"
	"boilerplate/internal/model"
	"boilerplate/internal/repository"
	"boilerplate/internal/usecase"
	"embed"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

//go:embed db/migrations/*.sql
var embedMigrations embed.FS

func main() {
	app := fiber.New()
	db := config.NewDB() //gorm
	db.Migrate(&embedMigrations)

	app.Get("/test", func(c *fiber.Ctx) error {
		res := model.ApiResponse[string]{}
		res.NotFound(c, "not found")
		return c.JSON(res)
	})

	useCase := usecase.NewUserUseCase(db.DB, validator.New(), repository.NewUserRepository())
	controller := http.NewUserController(useCase)

	app.Post("/login", controller.Login)

	secured := app.Group("secured")
	secured.Use(middleware.NewAuth())
	secured.Get("/hello", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})

	secured.Get("/hello/:name", func(ctx *fiber.Ctx) error {
		name := ctx.Params("name")
		return ctx.SendString("Hello " + name)
	})

	_ = app.Listen("localhost:3000")
}
