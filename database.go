package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func newDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%v?parseTime=true", *dbconn))
	if err != nil {
		fmt.Printf("error opening db: %v\n", err)
		return db, err
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("error pinging db: %v\n", err)
		return db, err
	}

	return db, nil
}

func GetDB() (*sql.DB, error) {
	var err error
	if db == nil {
		db, err = newDB()
	}
	return db, err
}
