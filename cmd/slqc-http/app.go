package main

import (
	"github.com/gadhittana01/go-sqlc/config"
	"github.com/gadhittana01/go-sqlc/handler/resthttp"
)

func initApp(c *config.GlobalConfig) error {

	return startHTTPServer(resthttp.NewRoutes(resthttp.RouterDependencies{}), c)
}
