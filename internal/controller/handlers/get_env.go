package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/marktsarkov/sigma-service/config"
)

func GetEnv(env *config.Environment) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Port: %v\nDatabase: %v", (*env).GetPort(), (*env).GetDB()))
	}
}
