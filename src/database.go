package main

import (
	"database/sql"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type databaseType struct {
	*sql.DB
	Driver           string
	ConnectionString string
	Password         string
}

var database databaseType

func (database *databaseType) Connect() error {
	connectionString := database.ConnectionString
	if database.Password != "" {
		connectionString = strings.Replace(connectionString, "@", ":"+database.Password+"@", 1)
		database.Password = ""
	}
	db, err := sql.Open(strings.ToLower(database.Driver), connectionString)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	database.DB = db

	database.DB.SetConnMaxLifetime(time.Minute * 3)
	database.DB.SetMaxOpenConns(10)
	database.DB.SetMaxIdleConns(10)

	return nil
}

func (database *databaseType) Query(query string) (*sql.Rows, error) {
	rows, err := database.DB.Query(query)
	if err != nil {
		pageAlert.show(err.Error(), "error")
	}

	return rows, err
}

func (database *databaseType) getTables() (*sql.Rows, error) {
	return database.Query("SHOW TABLES")
}
