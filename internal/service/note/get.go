package note

import (
	"context"
	"github.com/marktsarkov/sigma-service/internal/entity"
)

func (s *serv) GetById(ctx context.Context, id int64) (*entity.Note, error) {
	note, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return note, nil
}
