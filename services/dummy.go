package services

import (
	"context"
)

type DummyService interface {
	TestDummy(ctx context.Context) (string, error)
}

type dummyService struct {
	name    string
	isError bool
}

type DummyDependencies struct {
	name  string
	name2 string
}

func NewDummyService(name string, name2 DummyDependencies, isError bool) (DummyService, error) {
	return &dummyService{
		name:    name,
		isError: isError,
	}, nil
}

func (ar *dummyService) TestDummy(ctx context.Context) (string, error) {
	return "ok", nil
}
