package main

import (
	"log"

	"github.com/gadhittana01/go-sqlc/config"
	"github.com/gadhittana01/go-sqlc/helper"
)

func main() {
	config := &config.GlobalConfig{}
	helper.LoadConfig(config)
	err := initApp(config)
	if err != nil {
		log.Println(err)
	}
}
