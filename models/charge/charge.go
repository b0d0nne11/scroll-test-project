// Package charge models charge objects.
package charge

import (
	"time"

	"github.com/b0d0nne11/scroll-test-project/db"
	"github.com/b0d0nne11/scroll-test-project/models/account"
	"github.com/mailgun/log"
	"github.com/mailgun/scroll"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Charge represents a charge object for a specific account.
type Charge struct {
	ID        bson.ObjectId "_id"
	Account   mgo.DBRef
	Cents     int
	Timestamp time.Time
}

// New returns a new charge.
func New(account *account.Account, cents int, timestamp time.Time) *Charge {
	return &Charge{
		ID: bson.NewObjectId(),
		Account: mgo.DBRef{
			Collection: "account",
			Id:         account.ID,
			Database:   db.Get().Name,
		},
		Cents:     cents,
		Timestamp: timestamp,
	}
}

// Save persists a charge to the database.
func (c *Charge) Save() (*Charge, error) {
	collection := db.Get().C("charge")

	err := collection.Insert(c)
	if err != nil {
		log.Errorf("error creating %v: %v\n", c, err)
		return nil, err
	}

	return c, nil
}

func findBy(k string, v interface{}) (*Charge, error) {
	collection := db.Get().C("charge")

	var c Charge

	err := collection.Find(bson.M{k: v}).One(&c)
	if err == mgo.ErrNotFound {
		return nil, scroll.NotFoundError{Description: "not found"}
	}
	if err != nil {
		log.Errorf("error reading charge(%v=%v): %v\n", k, v, err)
		return nil, err
	}

	return &c, nil
}

// Get returns a charge with the given ID.
func Get(id string) (*Charge, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, scroll.InvalidFormatError{
			Field: "id",
			Value: "not an objectid",
		}
	}
	return findBy("_id", bson.ObjectIdHex(id))
}

// List returns a slice of `limit` charges skipping the first `last` charges.
func List(last int, limit int) ([]*Charge, error) {
	collection := db.Get().C("charge")

	var cl = make([]*Charge, 0, limit)

	err := collection.Find(nil).Skip(last).Limit(limit).All(&cl)
	if err != nil {
		log.Errorf("error listing charges(%v, %v)", last, limit)
		return nil, err
	}

	return cl, nil
}
