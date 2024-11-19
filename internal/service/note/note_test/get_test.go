package note_test

import (
	"context"
	"fmt"
	"github.com/marktsarkov/sigma-service/internal/entity"
	"github.com/marktsarkov/sigma-service/internal/repo/mocks"
	service "github.com/marktsarkov/sigma-service/internal/service/note"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	var (
		idFromHandler int64 = 1
		ctx                 = context.Background()
	)

	tc := []struct {
		name    string
		wantErr bool
		err     error
		id      int64
		note    *entity.Note
	}{
		{
			name:    "success case",
			wantErr: false,
			err:     nil,
			id:      idFromHandler,
			note: &entity.Note{
				ID:        1,
				Title:     "title",
				Body:      "body",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
		{
			name:    "fall case",
			wantErr: true,
			err:     fmt.Errorf("internal error"),
			id:      idFromHandler,
			note: &entity.Note{
				ID:        1,
				Title:     "title",
				Body:      "body",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {

			repoMock := new(mocks.NoteRepository)
			repoMock.EXPECT().GetById(mock.Anything, idFromHandler).Return(tt.note, tt.err)

			serv := service.NewNoteService(repoMock)
			returnedNote, err := serv.GetById(ctx, tt.id)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.note, returnedNote)
			}
			repoMock.AssertExpectations(t)

		})
	}
}
