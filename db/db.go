//  Packge db initializes and maintains database state information
package db

import (
	"github.com/mailgun/log"
	"gopkg.in/mgo.v2"
)

var dbh *mgo.Database

// DatabaseConfig contains database connection and configuration parameters
type DatabaseConfig struct {
	User     string `config:"optional"`
	Password string `config:"optional"`
	Address  string
	DB       string
}

// Init initializes a database session and returns a session object. This
// session object should be closed when the calling function finishes with it. It
// has the side effect of saving the database object that is returned by the Get
// function.
//
// Example:
//     dbSession, err := db.Init(dbConfig)
//     if err != nil {
//     	 panic(fmt.Sprintf("error connecting to db: %v\n", err))
//     }
//     defer dbSession.Close()
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

// Get returns a database object. It should only be called after Init.
func Get() *mgo.Database {
	return dbh
}
