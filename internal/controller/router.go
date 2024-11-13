package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marktsarkov/sigma-service/config"
	"github.com/marktsarkov/sigma-service/internal/controller/handlers"
	"github.com/marktsarkov/sigma-service/internal/service"
)

func NewRouter(app *fiber.App, env *config.Environment, service service.NoteService) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/env", handlers.GetEnv(env))
	app.Post("/notes", handlers.CreateNote(service))
	app.Get("/notes", handlers.GetNoteById(service))
}
