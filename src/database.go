package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type databaseType struct {
	*sql.DB
	Driver           string
	DriverName       string
	Host             string
	User             string
	Password         string
	Database         string
	ConnectionString string
}

var database databaseType

func (database *databaseType) buildConnectionString() {
	if database.DriverName == "MySQL/MariaDB" {
		database.Driver = "mysql"
		database.ConnectionString = database.User + ":" + database.Password + "@tcp(" + database.Host + ")/" + database.Database
	} else if database.DriverName == "PostgreSQL" {
		database.Driver = "postgres"
		database.ConnectionString = "host=" + database.Host + " user=" + database.User + " password=" + database.Password + " dbname=" + database.Database + " sslmode=disable"
	}
}

func (database *databaseType) Connect() error {
	database.buildConnectionString()
	db, err := sql.Open(database.Driver, database.ConnectionString)
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
	query := ""
	if database.Driver == "mysql" {
		query = "SHOW TABLES"
	} else if database.Driver == "postgres" {
		query = "SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname != 'pg_catalog' AND schemaname != 'information_schema'"
	}
	return database.Query(query)
}
