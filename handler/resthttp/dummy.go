package resthttp

import (
	"context"
	"net/http"
)

type DummyHandler struct {
	dummyService DummyService
}

type DummyDependencies struct {
	name string
}

func NewDummyHandler(dep DummyDependencies, dummyService DummyService) *DummyHandler {
	return &DummyHandler{}
}

func (p DummyHandler) TestDummy(w http.ResponseWriter, r *http.Request) {
	resp := NewResponse()

	res, err := p.dummyService.TestDummy(context.Background())
	if err != nil {
		resp.SetInternalServerError(err.Error(), w)
		return
	}

	resp.SetOK(res, w)
	return
}
