package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marktsarkov/sigma-service/config"
	handlers2 "github.com/marktsarkov/sigma-service/internal/controller/http/handlers"
	"github.com/marktsarkov/sigma-service/internal/service"
)

func NewRouter(app *fiber.App, env *config.Environment, service service.NoteService) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/env", handlers2.GetEnv(env))
	app.Post("/notes", handlers2.CreateNote(service))
	app.Get("/notes", handlers2.GetNoteById(service))
}
