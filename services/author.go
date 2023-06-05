package services

import (
	"context"
)

type AuthorService interface {
	CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error)
	DeleteAuthor(ctx context.Context, id int64) error
	GetAuthor(ctx context.Context, id int64) (Author, error)
	ListAuthors(ctx context.Context) ([]Author, error)
	UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) (Author, error)
}

type authorService struct {
	ar AuthorResource
}

func NewAuthorService(AR AuthorResource) (AuthorService, error) {
	return &authorService{
		ar: AR,
	}, nil
}

func (ar *authorService) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error) {
	var res Author = Author{}
	return res, nil
}
func (ar *authorService) DeleteAuthor(ctx context.Context, id int64) error {
	return nil
}
func (ar *authorService) GetAuthor(ctx context.Context, id int64) (Author, error) {
	var res Author = Author{}
	return res, nil
}
func (ar *authorService) ListAuthors(ctx context.Context) ([]Author, error) {
	var res []Author = []Author{}
	return res, nil
}
func (ar *authorService) UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) (Author, error) {
	var res Author = Author{}
	return res, nil
}
