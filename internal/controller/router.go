package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marktsarkov/sigma-service/config"
	"github.com/marktsarkov/sigma-service/internal/controller/handlers"
	repository "github.com/marktsarkov/sigma-service/internal/repo"
)

func NewRouter(app *fiber.App, env *config.Environment, noteRepo repository.NoteRepository) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/env", handlers.GetEnv(env))
	app.Post("/notes", handlers.CreateNote(noteRepo))
}
