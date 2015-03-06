package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mailgun/log"
)

var dbh *sql.DB

type DatabaseConfig struct {
	User     string
	Password string
	Address  string `config:"optional"`
	DB       string
}

func conn(conf DatabaseConfig) string {
	return fmt.Sprintf("%v:%v@%v/%v?parseTime=true", conf.User, conf.Password, conf.Address, conf.DB)
}

func Init(conf DatabaseConfig) (*sql.DB, error) {
	var err error
	dbh, err = sql.Open("mysql", conn(conf))
	if err != nil {
		log.Errorf("error opening db: %v\n", err)
		dbh = nil
		return nil, err
	}
	err = dbh.Ping()
	if err != nil {
		log.Errorf("error pinging db: %v\n", err)
		dbh = nil
		return nil, err
	}
	return dbh, nil
}

func Get() *sql.DB {
	return dbh
}
