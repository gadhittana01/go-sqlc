package resthttp

import (
	"context"

	"github.com/gadhittana01/go-sqlc/services"
)

type (
	AuthorService interface {
		CreateAuthor(ctx context.Context, arg services.CreateAuthorParams) (services.Author, error)
		DeleteAuthor(ctx context.Context, id int64) error
		GetAuthor(ctx context.Context, id int64) (services.Author, error)
		ListAuthors(ctx context.Context) ([]services.Author, error)
		UpdateAuthor(ctx context.Context, arg services.UpdateAuthorParams) (services.Author, error)
	}

	DummyService interface {
		TestDummy(ctx context.Context) (string, error)
	}
)
