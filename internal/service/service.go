package service

import (
	"context"
	"github.com/marktsarkov/sigma-service/internal/entity"
)

type NoteService interface {
	Create(ctx context.Context, note *entity.Note) (int64, error)
	GetById(ctx context.Context, id int64) (*entity.Note, error)
}
