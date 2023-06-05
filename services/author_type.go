package services

import "database/sql"

type CreateAuthorParams struct {
	Name string
	Bio  sql.NullString
}

type UpdateAuthorParams struct {
	ID   int64
	Name string
	Bio  sql.NullString
}

type Author struct {
	ID   int64
	Name string
	Bio  sql.NullString
}

type AuthorDependencies struct {
	AR AuthorResource
}
