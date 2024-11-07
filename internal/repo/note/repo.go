package note

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/marktsarkov/sigma-service/internal/entity"
	"github.com/marktsarkov/sigma-service/internal/repo"
	"github.com/marktsarkov/sigma-service/internal/repo/note/converter"
	"github.com/marktsarkov/sigma-service/internal/repo/note/model"
)

const (
	tableName       = "note"
	idColumn        = "id"
	titleColumn     = "title"
	bodyColumn      = "body"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

//пишем функции с sql-запросами в бд и отправляем их, возвращаем результат

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.NoteRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, note *entity.Note) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(titleColumn, bodyColumn).
		Values(note.Title, note.Body).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}
	var id int64

	err = r.db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repo) GetById(ctx context.Context, id int64) (*entity.Note, error) {
	builder := sq.Select(idColumn, titleColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{"id": id})
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var note model.Note

	err = r.db.QueryRow(ctx, query, args...).Scan(&note.ID, &note.Title, &note.Body, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.ToNoteFromRepo(&note), nil
}
