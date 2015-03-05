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

func setup() {
	dbh, err = sql.Open("mysql", *dbconn+"?parseTime=true")
	if err != nil {
		fmt.Printf("error opening db: %v\n", err)
		return
	}
	err = dbh.Ping()
	if err != nil {
		fmt.Printf("error pinging db: %v\n", err)
		return
	}
}

func Get() (*sql.DB, error) {
	if dbh == nil {
		setup()
	}
	if err != nil {
		return nil, err
	}
	return dbh, nil
}
