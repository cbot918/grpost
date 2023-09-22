package server

import (
	db "github.com/cbot918/grpost/db/sqlc"
	"github.com/cbot918/grpost/server/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type App struct {
	Server *fiber.App
	Query  *db.Queries
}

func New(uiPath string, query *db.Queries) *App {

	app := new(App)

	app.Server = fiber.New()
	app.Server.Use(cors.New()) // setup allow cors *
	app.Server.Use(recover.New())
	app.Server.Use(logger.New())

	app.Server.Static("/", uiPath) // serve spa

	// api routing
	app.Server.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"code": 200, "message": "Hello, World"})
	})

	c := controller.NewController(query)
	app.Server.Post("/signin", c.Auth.Signin)
	app.Server.Post("/signup", c.Auth.Signup)
	app.Server.Get("/allpost", c.Post.AllPost)

	return app
}
