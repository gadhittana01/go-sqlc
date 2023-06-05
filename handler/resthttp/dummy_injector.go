//go:build wireinject
// +build wireinject

package resthttp

import (
	"github.com/gadhittana01/go-sqlc/services"
	"github.com/google/wire"
)

var dummyService = wire.NewSet(
	services.NewDummyService,
	wire.Bind(new(DummyService), new(services.DummyService)),
)

func InitializedDummyHandler(name string, name2 services.DummyDependencies, isError bool, dep DummyDependencies) (*DummyHandler, error) {
	wire.Build(dummyService, NewDummyHandler)
	return nil, nil
}
