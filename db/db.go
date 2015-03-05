package db

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var dbh *sql.DB
var err error
var dbconn *string = flag.String("db", "username:password@address/database", "Database connection string")

func setup() (*sql.DB, error) {
	dbh, err = sql.Open("mysql", *dbconn+"?parseTime=true")
	if err != nil {
		fmt.Printf("error opening db: %v\n", err)
		return dbh, err
	}
	err = dbh.Ping()
	if err != nil {
		fmt.Printf("error pinging db: %v\n", err)
		return dbh, err
	}

	return dbh, nil
}

func Get() (*sql.DB, error) {
	if dbh == nil {
		dbh, err = setup()
	}
	return dbh, err
}
