package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ruegerj/stock-sight/internal/common"
	"github.com/ruegerj/stock-sight/internal/db"
	"github.com/ruegerj/stock-sight/internal/queries"
)

var empty = queries.Author{}

type AuthorRepository interface {
	GetById(ctx context.Context, id int64) (queries.Author, error)
	GetAll(ctx context.Context) ([]queries.Author, error)
	Create(ctx context.Context, name string, bio *string) (queries.Author, error)
	Update(ctx context.Context, id int64, name string, bio *string) error
	Delete(ctx context.Context, id int64) error
}

func NewSqlcAuthorRepository(connection db.DbConnection) AuthorRepository {
	return &SqlcAuthorRepository{
		queries: queries.New(connection.Database()),
	}
}

type SqlcAuthorRepository struct {
	queries *queries.Queries
}

func (sar *SqlcAuthorRepository) GetAll(ctx context.Context) ([]queries.Author, error) {
	return sar.queries.ListAuthors(ctx)
}

func (sar *SqlcAuthorRepository) GetById(ctx context.Context, id int64) (queries.Author, error) {
	if id < 0 {
		return empty, errors.New("Supply a valid author id")
	}

	return sar.queries.GetAuthor(ctx, id)
}

func (sar *SqlcAuthorRepository) Create(ctx context.Context, name string, bio *string) (queries.Author, error) {
	if len(name) <= 0 {
		return empty, errors.New("Name must contain at least one character")
	}

	createParams := queries.CreateAuthorParams{
		Name: name,
		Bio: sql.NullString{
			String: common.DerefOrEmpty(bio),
			Valid:  bio != nil,
		},
	}

	return sar.queries.CreateAuthor(ctx, createParams)
}

func (sar *SqlcAuthorRepository) Update(ctx context.Context, id int64, name string, bio *string) error {
	if id < 0 {
		return errors.New("Supply a valid author id")
	}
	if len(name) <= 0 {
		return errors.New("Name must contain at least one character")
	}

	updateParams := queries.UpdateAuthorParams{
		ID:   id,
		Name: name,
		Bio: sql.NullString{
			String: common.DerefOrEmpty(bio),
			Valid:  bio != nil,
		},
	}

	return sar.queries.UpdateAuthor(ctx, updateParams)
}

func (sar *SqlcAuthorRepository) Delete(ctx context.Context, id int64) error {
	if id < 0 {
		return errors.New("Supply a valid author id")
	}

	return sar.queries.DeleteAuthor(ctx, id)
}
