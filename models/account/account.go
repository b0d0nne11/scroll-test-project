package account

import (
	"github.com/b0d0nne11/scroll-test-project/db"
	"github.com/mailgun/log"
	"github.com/mailgun/scroll"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	ID   bson.ObjectId "_id"
	Name string
}

func New(name string) *Account {
	return &Account{
		ID:   bson.NewObjectId(),
		Name: name,
	}
}

func (a *Account) Save() (*Account, error) {
	collection := db.Get().C("account")

	err := collection.Insert(a)
	if err != nil {
		log.Errorf("error creating %v: %v\n", a, err)
		return nil, err
	}

	return a, nil
}

func findBy(k string, v interface{}) (*Account, error) {
	collection := db.Get().C("account")

	var a Account

	err := collection.Find(bson.M{k: v}).One(&a)
	if err == mgo.ErrNotFound {
		return nil, scroll.NotFoundError{Description: "not found"}
	}
	if err != nil {
		log.Errorf("error reading account(%v=%v): %v\n", k, v, err)
		return nil, err
	}

	return &a, nil
}

func Get(id string) (*Account, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, scroll.InvalidFormatError{
			Field: "id",
			Value: "not an objectid",
		}
	}
	return findBy("_id", bson.ObjectIdHex(id))
}

func FindByName(name string) (*Account, error) {
	return findBy("name", name)
}

func List(last int, limit int) ([]*Account, error) {
	collection := db.Get().C("account")

	var al = make([]*Account, 0, limit)

	err := collection.Find(nil).Skip(last).Limit(limit).All(&al)
	if err != nil {
		log.Errorf("error listing accounts(%v, %v)", last, limit)
		return nil, err
	}

	return al, nil
}
