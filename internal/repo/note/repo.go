package note

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/marktsarkov/sigma-service/internal/entity"
	"github.com/marktsarkov/sigma-service/internal/repo"
	"github.com/marktsarkov/sigma-service/internal/repo/note/converter"
	"github.com/marktsarkov/sigma-service/internal/repo/note/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const (
	tableName       = "note"
	idColumn        = "id"
	titleColumn     = "title"
	bodyColumn      = "body"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	pg *pgxpool.Pool
	mg *mongo.Client
}

func NewRepository(pg *pgxpool.Pool, mg *mongo.Client) repository.NoteRepository {
	return &repo{
		pg: pg,
		mg: mg,
	}
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

	err = r.pg.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}
	note.ID = id
	//добавить замтеку в монгу с ид сгенерированным в постгре
	collection := r.mg.Database("Notes").Collection("notes")

	mgNote, err := bson.Marshal(note)
	resultMG, err := collection.InsertOne(ctx, mgNote)
	if err != nil {
		return 0, err
	}
	log.Printf("В монгу поступила заметка: %d\n", resultMG.InsertedID)
	return id, nil
}

func (r *repo) GetById(ctx context.Context, id int64) (*entity.Note, error) {
	builder := sq.Select(idColumn, titleColumn, bodyColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{"id": id})
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var note model.Note

	err = r.pg.QueryRow(ctx, query, args...).Scan(&note.ID, &note.Title, &note.Body, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		return nil, err
	}

	collection := r.mg.Database("Notes").Collection("notes")

	err = collection.FindOne(ctx, bson.M{"id": note.ID}).Decode(&note)
	if err != nil {
		return nil, err
	}
	log.Printf("Из монги пришла заметка: %d\n", &note.Body)

	return converter.ToNoteFromRepo(&note), nil
}
