package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/marktsarkov/sigma-service/internal/config"
)

type environment struct {
	Port int `envconfig:"PORT"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	env, err := config.NewEnv()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/env", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Port: %v\nDatabase: %v", env.GetPort(), env.GetDB()))
	})

	app.Listen(fmt.Sprintf(":%v", env.GetPort()))
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
