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
		idFromHandler int64        = 1
		expectedNote  *entity.Note = &entity.Note{
			ID:        1,
			Title:     "title",
			Body:      "body",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		ctx = context.Background()
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
			note:    expectedNote,
		},
		{
			name:    "fall case",
			wantErr: true,
			err:     fmt.Errorf("internal error"),
			id:      idFromHandler,
			note:    expectedNote,
		},
	}

	for _, tt := range tc {

		repoMock := new(mocks.NoteRepository)
		repoMock.EXPECT().GetById(mock.Anything, idFromHandler).Return(expectedNote, nil)

		serv := service.NewNoteService(repoMock)

		returnedNote, err := serv.GetById(ctx, tt.id)

		require.NoError(t, err)
		assert.Equal(t, expectedNote, returnedNote)
		repoMock.AssertExpectations(t)
	}
}
