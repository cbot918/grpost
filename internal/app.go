package internal

import (
	"github.com/cbot918/grpost/internal/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type App struct {
	Server *fiber.App
	port   string
}

func New(port string, uiPath string) *App {

	app := new(App)

	app.Server = fiber.New()

	app.Server.Use(cors.New())     // setup use cors
	app.Server.Static("/", uiPath) // serve spa
	app.port = port

	return app
}

func (a *App) Run() {

	a.setupRouting()

	a.Server.Listen(a.port)
}

func (a *App) setupRouting() {
	a.Server.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"code": 200, "message": "Hello, World"})
	})

	c := controller.NewController()

	a.Server.Post("/signin", c.Auth.Signin)
}
