package main

import (
	"boilerplate/internal/config"
	"boilerplate/internal/http"
	"boilerplate/internal/middleware"
	"boilerplate/internal/model"
	"boilerplate/internal/repository"
	"boilerplate/internal/usecase"
	"embed"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

//go:embed db/migrations/*.sql
var embedMigrations embed.FS

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			switch code {
			case 500:
				return ctx.JSON((&model.ApiResponse[any]{}).ServerError(e.Message))
			}

			return ctx.Status(code).JSON("Server Error")
		},
	})
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

	log.Fatal(app.Listen("localhost:3000"))
}
