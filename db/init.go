package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gadhittana01/go-sqlc/config"
	"github.com/gadhittana01/go-sqlc/helper"
	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	var db *sql.DB
	var err error

	config := &config.GlobalConfig{}
	helper.LoadConfig(config)

	dbConn := config.DB

	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConn.Host, dbConn.Port, dbConn.User, dbConn.Password, dbConn.Name,
	)

	db, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB " + dbConn.Name + " connected Successfully!")

	return db
}
