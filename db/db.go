package db

import (
	"github.com/mailgun/log"
	"gopkg.in/mgo.v2"
)

var dbh *mgo.Database

type DatabaseConfig struct {
	User     string `config:"optional"`
	Password string `config:"optional"`
	Address  string
	DB       string
}

func Init(conf DatabaseConfig) (*mgo.Session, error) {
	session, err := mgo.Dial(conf.Address)
	if err != nil {
		log.Errorf("error opening db session: %v\n", err)
		return nil, err
	}
	err = session.Ping()
	if err != nil {
		log.Errorf("error pinging db session: %v\n", err)
		return nil, err
	}
	dbh = session.DB(conf.DB)
	if conf.User != "" {
		err = dbh.Login(conf.User, conf.Password)
		if err != nil {
			log.Errorf("error logging into db: %v", err)
			dbh = nil
			return nil, err
		}
	}

	return session, nil
}

func Get() *mgo.Database {
	return dbh
}
