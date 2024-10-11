package config

import (
	"embed"
	"github.com/gofiber/fiber/v2/log"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBInstance struct {
	DB *gorm.DB
}

func NewDB() DBInstance {
	dsn := "host=localhost user=postgres password=postgres dbname=gofiber_boilerplate port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed connect database: %v", err)
	}

	return DBInstance{DB: db}
}

func (db *DBInstance) Migrate(embedMigrations *embed.FS) {
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
}
