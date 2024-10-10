package main

import (
	"boilerplate/internal/config"
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/pressly/goose/v3"
	"log"
)

//go:embed db/migrations/*.sql
var embedMigrations embed.FS

func main() {
	app := fiber.New()
	db := config.NewDB() //gorm
	s, err := db.DB.DB()
	if err != nil {
		log.Fatalf("failed to get *sql.DB from gorm: %v", err)
	}
	sqlDB := s // *sql.DB

	// goose embed migrations
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("failed to set dialect: %v", err)
	}
	if err := goose.Up(sqlDB, "db/migrations"); err != nil {
		log.Fatalf("migrations failed: %v", err)
	}

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})

	app.Get("/:name", func(ctx *fiber.Ctx) error {
		name := ctx.Params("name")
		return ctx.SendString("Hello " + name)
	})

	_ = app.Listen("localhost:3000")
}
