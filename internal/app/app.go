package app

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/marktsarkov/sigma-service/internal/repo/note"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/marktsarkov/sigma-service/config"
	"github.com/marktsarkov/sigma-service/internal/controller"
)

func Run(cfg *config.Environment, ctx context.Context) {
	app := fiber.New()

	DSN := os.Getenv("DATABASE_URL")
	pool, err := pgxpool.New(ctx, DSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	repo := note.NewRepository(pool)
	controller.NewRouter(app, cfg, repo)

	if err := app.Listen(fmt.Sprintf(":%v", (*cfg).GetPort())); err != nil {
		log.Fatal(err)
	}
}
