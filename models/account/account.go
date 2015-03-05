package account

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/b0d0nne11/scroll-test-project/db"
	"github.com/mailgun/scroll"
)

type Account struct {
	ID   int64
	Name string
}

func New(name string) *Account {
	return &Account{
		Name: name,
	}
}

func (a *Account) Save() error {
	dbh, _ := db.Get()

	stmt, err := dbh.Prepare("INSERT INTO account (name) VALUES ( ? )")
	if err != nil {
		fmt.Printf("error preparing statement: %v\n", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(a.Name)
	if err != nil {
		fmt.Printf("error creating %v: %v\n", a, err)
		return err
	}
	a.ID, err = res.LastInsertId()
	if err != nil {
		fmt.Printf("error creating %v: %v\n", a, err)
		return err
	}

	return nil
}

func GetBy(k string, v string) (*Account, error) {
	dbh, _ := db.Get()

	var a Account

	stmt, err := dbh.Prepare(fmt.Sprintf("SELECT id, name FROM account WHERE %v = ?", k))
	if err != nil {
		fmt.Printf("error preparing statement: %v\n", err)
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(v).Scan(&a.ID, &a.Name)
	if err == sql.ErrNoRows {
		return nil, scroll.NotFoundError{
			Description: "account not found",
		}
	}
	if err != nil {
		fmt.Printf("error reading account(%v=%v): %v\n", k, v, err)
		return nil, err
	}

	return &a, nil
}

func Get(id int64) (*Account, error) {
	return GetBy("id", strconv.FormatInt(id, 10))
}

func List(last int64, limit int64) ([]*Account, error) {
	dbh, _ := db.Get()

	var al = make([]*Account, 0, limit)

	stmt, err := dbh.Prepare("SELECT id, name FROM account WHERE id > ? LIMIT ?")
	if err != nil {
		fmt.Printf("error preparing statement: %v\n", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(last, limit)
	if err != nil {
		fmt.Printf("error listing accounts(%v, %v)", last, limit)
		return nil, err
	}

	for rows.Next() {
		var a Account
		err = rows.Scan(&a.ID, &a.Name)
		if err != nil {
			fmt.Printf("error listing accounts(%v, %v)", last, limit)
			return nil, err
		}
		al = append(al, &a)
	}

	return al, nil
}
