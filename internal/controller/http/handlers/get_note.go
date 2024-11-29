package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/marktsarkov/sigma-service/internal/entity"
	"github.com/marktsarkov/sigma-service/internal/service"
	"strconv"
	"time"
)

func GetNoteById(service service.NoteService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var note *entity.Note

		idParam := c.Params("id")
		id, err := strconv.ParseInt(idParam, 10, 64)

		if err != nil {
			fmt.Println(err, "parser")
			c.Status(fiber.StatusBadRequest).SendString(err.Error()) // Возвращаем ошибку, если не смогли распарсить тело запроса
			return err

		}

		note, err = service.GetById(c.Context(), id)
		if err != nil {
			fmt.Println(err, "to note")
			c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // Возвращаем ошибку сервера, если что-то пошло не так
			return err
		}
		response := map[string]string{
			"id":         strconv.FormatInt(id, 10),
			"title":      note.Title,
			"body":       note.Body,
			"created_at": note.CreatedAt.Format(time.RFC3339),
			"updated_at": note.UpdatedAt.Format(time.RFC3339),
		}

		return c.JSON(response)
	}
}
