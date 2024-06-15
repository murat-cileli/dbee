package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type databaseType struct {
	*sql.DB
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func (database *databaseType) Connect() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=10s", database.User, database.Password, database.Host, database.Port, database.Database))
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
