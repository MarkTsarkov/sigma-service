package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/marktsarkov/sigma-service/config"
	"github.com/marktsarkov/sigma-service/internal/app"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	env, err := config.NewEnv()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	app.Run(&env, ctx)
}
