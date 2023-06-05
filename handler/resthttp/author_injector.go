//go:build wireinject
// +build wireinject

package resthttp

import (
	"database/sql"

	"github.com/gadhittana01/go-sqlc/db"
	"github.com/gadhittana01/go-sqlc/pkg/author"
	"github.com/gadhittana01/go-sqlc/services"
	"github.com/google/wire"
)

var dbSet = wire.NewSet(
	db.InitDB,
	wire.Bind(new(author.DBTX), new(*sql.DB)),
)

var authorPkgSet = wire.NewSet(
	author.New,
	wire.Bind(new(services.AuthorResource), new(*author.Queries)),
)

var authorService = wire.NewSet(
	services.NewAuthorService,
	wire.Bind(new(AuthorService), new(services.AuthorService)),
)

func InitializedAuthorHandler() (*AuthorHandler, error) {
	wire.Build(dbSet, authorPkgSet, authorService, NewAuthorHandler)
	return nil, nil
}
