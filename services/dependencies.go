package services

import (
	"context"

	"github.com/gadhittana01/go-sqlc/pkg/author"
)

type (
	AuthorResource interface {
		CreateAuthor(ctx context.Context, arg author.CreateAuthorParams) (author.Author, error)
		DeleteAuthor(ctx context.Context, id int64) error
		GetAuthor(ctx context.Context, id int64) (author.Author, error)
		ListAuthors(ctx context.Context) ([]author.Author, error)
		UpdateAuthor(ctx context.Context, arg author.UpdateAuthorParams) (author.Author, error)
	}
)
