package resthttp

import (
	"context"
	"net/http"
)

type AuthorHandler struct {
	authorService AuthorService
}

func NewAuthorHandler(authorService AuthorService) *AuthorHandler {
	return &AuthorHandler{
		authorService: authorService,
	}
}

func (p AuthorHandler) GetListAuthor(w http.ResponseWriter, r *http.Request) {
	resp := NewResponse()

	res, err := p.authorService.ListAuthors(context.Background())
	if err != nil {
		resp.SetInternalServerError(err.Error(), w)
		return
	}

	resp.SetOK(res, w)
	return
}
