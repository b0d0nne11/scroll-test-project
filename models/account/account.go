package account

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/mailgun/scroll"
)

type Account struct {
	ID   int64
	Name string
}

func New(name string) Account {
	return Account{
		Name: name,
	}
}

func (a *Account) Save(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO account (name) VALUES ( ? )")
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

func GetBy(db *sql.DB, k string, v string) (Account, error) {
	var a Account

	stmt, err := db.Prepare(fmt.Sprintf("SELECT id, name FROM account WHERE %v = ?", k))
	if err != nil {
		fmt.Printf("error preparing statement: %v\n", err)
		return a, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(v).Scan(&a.ID, &a.Name)
	if err == sql.ErrNoRows {
		return a, scroll.NotFoundError{
			Description: "account not found",
		}
	}
	if err != nil {
		fmt.Printf("error reading account(%v=%v): %v\n", k, v, err)
		return a, err
	}

	return a, nil
}

func Get(db *sql.DB, id int64) (Account, error) {
	return GetBy(db, "id", strconv.FormatInt(id, 10))
}

func List(db *sql.DB, last int64, limit int64) ([]Account, error) {
	var al = make([]Account, 0, limit)

	stmt, err := db.Prepare("SELECT id, name FROM account WHERE id > ? LIMIT ?")
	if err != nil {
		fmt.Printf("error preparing statement: %v\n", err)
		return al, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(last, limit)
	if err != nil {
		fmt.Printf("error listing accounts(%v, %v)", last, limit)
		return al, err
	}

	for rows.Next() {
		var a Account
		err = rows.Scan(&a.ID, &a.Name)
		if err != nil {
			fmt.Printf("error listing accounts(%v, %v)", last, limit)
			return al, err
		}
		al = append(al, a)
	}

	return al, nil
}
