package handlers

import (
	"context"
	pb "github.com/marktsarkov/sigma-service/internal/controller/grpc"
	"github.com/marktsarkov/sigma-service/internal/entity"
)

func (s pb.NoteServer) GetByID(ctx context.Context, id int64) (note *entity.Note, err error) {

}
