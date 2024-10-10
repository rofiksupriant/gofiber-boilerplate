package config

import (
	"github.com/gofiber/fiber/v2/log"
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
		log.Error("failed connect database")
		panic(err)
	}

	return DBInstance{DB: db}
}
