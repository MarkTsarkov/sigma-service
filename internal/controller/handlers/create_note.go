package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marktsarkov/sigma-service/internal/entity"
	repository "github.com/marktsarkov/sigma-service/internal/repo"
	"strconv"
)

func CreateNote(repo repository.NoteRepository) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var body entity.Note

		err := c.BodyParser(&body)
		if err != nil {
			c.Status(fiber.StatusBadRequest).SendString(err.Error()) // Возвращаем ошибку, если не смогли распарсить тело запроса
			return err
		}

		id, err := repo.Create(c.Context(), &body)
		if err != nil {
			c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // Возвращаем ошибку сервера, если что-то пошло не так
			return err
		}

		response := map[string]string{"id": strconv.FormatInt(id, 10)}
		return c.JSON(response)
	}
}
