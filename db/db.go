package db

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mailgun/log"
)

var dbh *sql.DB
var err error
var dbconn *string = flag.String("db", "username:password@address/database", "Database connection string")

func setup() {
	dbh, err = sql.Open("mysql", *dbconn+"?parseTime=true")
	if err != nil {
		log.Errorf("error opening db: %v\n", err)
		return
	}
	err = dbh.Ping()
	if err != nil {
		log.Errorf("error pinging db: %v\n", err)
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
