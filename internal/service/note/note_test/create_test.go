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
		ctx = context.Background()
	)

	tc := []struct {
		name       string
		wantErr    bool
		mockReturn struct {
			id  int64
			err error
		}
		note *entity.Note
	}{
		{
			name:    "success case",
			wantErr: false,
			mockReturn: struct {
				id  int64
				err error
			}{1, nil},
			note: &entity.Note{Title: "title", Body: "body"},
		},
		{
			name:    "fall case",
			wantErr: true,
			mockReturn: struct {
				id  int64
				err error
			}{0, fmt.Errorf("internal error")},
			note: &entity.Note{Title: "title", Body: "body"},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			repoMock := new(mocks.NoteRepository)
			repoMock.EXPECT().Create(mock.Anything, tt.note).Return(tt.mockReturn.id, tt.mockReturn.err)

			serv := service.NewNoteService(repoMock)

			id, err := serv.Create(ctx, tt.note)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.mockReturn.id, id)
			}

			repoMock.AssertExpectations(t)
		})
	}
}
