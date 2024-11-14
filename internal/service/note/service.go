package note

import (
	repository "github.com/marktsarkov/sigma-service/internal/repo"
	"github.com/marktsarkov/sigma-service/internal/service"
)

type serv struct {
	repo repository.NoteRepository
}

func NewNoteService(repo repository.NoteRepository) service.NoteService {
	return &serv{repo: repo}
}
