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

func (database *databaseType) Connect() error {
	if database.Password != "" {
		database.ConnectionString = strings.Replace(database.ConnectionString, "@", ":"+database.Password+"@", 1)
	}
	db, err := sql.Open(strings.ToLower(database.Driver), database.ConnectionString)
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

func (database *databaseType) Query(query string) *sql.Rows {
	rows, err := database.DB.Query(query)
	if err != nil {
		pageAlert.show(err.Error(), "error")
		return nil
	}

	return rows
}

func (database *databaseType) getTables() *sql.Rows {
	return database.Query("SHOW TABLES")
}
