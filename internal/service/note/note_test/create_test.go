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
)

func TestCreate(t *testing.T) {
	var (
		idFromDB int64        = 1
		note     *entity.Note = &entity.Note{Title: "title", Body: "body"}
		err      error
		ctx      = context.Background()
	)

	tc := []struct {
		name    string
		wantErr bool
		err     error
		note    *entity.Note
	}{
		{
			name:    "success case",
			wantErr: false,
			err:     nil,
			note:    note,
		},
		{
			name:    "fall case",
			wantErr: true,
			err:     fmt.Errorf("internal error"),
			note:    note,
		},
	}

	for _, tt := range tc {

		repoMock := new(mocks.NoteRepository)
		repoMock.EXPECT().Create(mock.Anything, note).Return(idFromDB, nil)

		serv := service.NewNoteService(repoMock)

		note.ID, err = serv.Create(ctx, tt.note)

		require.NoError(t, err)
		assert.Equal(t, idFromDB, note.ID)
		repoMock.AssertExpectations(t)
	}
}
