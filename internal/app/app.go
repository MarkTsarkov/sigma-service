package app

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/marktsarkov/sigma-service/internal/controller/http"
	repository "github.com/marktsarkov/sigma-service/internal/repo/note"
	service "github.com/marktsarkov/sigma-service/internal/service/note"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/marktsarkov/sigma-service/config"
)

func Run(cfg *config.Environment, ctx context.Context) {
	app := fiber.New()
	//pg
	DSN := os.Getenv("DATABASE_URL")
	pool, err := pgxpool.New(ctx, DSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	//mongo
	URI := os.Getenv("MONGO_URI")
	mg, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := mg.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	// Проверка подключения
	if err := mg.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Подключение успешно!")

	repo := repository.NewRepository(pool, mg)
	serv := service.NewNoteService(repo)

	var wg sync.WaitGroup
	wg.Add(1)
	//http :8081
	go func() {
		defer wg.Done()
		http.NewRouter(app, cfg, serv)
		if err := app.Listen(fmt.Sprintf(":%v", (*cfg).GetPort())); err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()
}
