package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

const dsn = "postgres://postgres:12345@localhost:5432/grpost?sslmode=disable"

func main() {

	api := fiber.New()

	api.Use(cors.New())
	api.Use(recover.New())
	api.Use(logger.New())

	api.Listen("")

}
