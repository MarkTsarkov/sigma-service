package app

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	desc "github.com/marktsarkov/sigma-service/internal/controller/grpc"
	"github.com/marktsarkov/sigma-service/internal/controller/http"
	repository "github.com/marktsarkov/sigma-service/internal/repo/note"
	service "github.com/marktsarkov/sigma-service/internal/service/note"
	pb "github.com/marktsarkov/sigma-service/pkg/note"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/marktsarkov/sigma-service/config"
)

func Run(cfg *config.Environment, ctx context.Context) {
	app := fiber.New()

	DSN := os.Getenv("DATABASE_URL")
	pool, err := pgxpool.New(ctx, DSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	repo := repository.NewRepository(pool)
	serv := service.NewNoteService(repo)

	var wg sync.WaitGroup
	wg.Add(2)
	//http :8081
	go func() {
		defer wg.Done()
		http.NewRouter(app, cfg, serv)
		if err := app.Listen(fmt.Sprintf(":%v", (*cfg).GetPort())); err != nil {
			log.Fatal(err)
		}
	}()

	//grpc :8082
	go func() {
		defer wg.Done()
		lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8082))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		var opts []grpc.ServerOption
		grpcServer := grpc.NewServer(opts...)
		reflection.Register(grpcServer)
		pb.RegisterNoteServer(grpcServer, desc.NewNoteServer(pb.UnimplementedNoteServer{}, serv))
		grpcServer.Serve(lis)
	}()

	wg.Wait()
}
