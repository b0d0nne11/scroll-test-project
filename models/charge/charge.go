package charge

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/mailgun/scroll"
)

type Charge struct {
	ID        int64
	AccountID int64
	Cents     uint64
	Timestamp time.Time
}

func New(accountID int64, cents uint64, timestamp time.Time) Charge {
	return Charge{
		AccountID: accountID,
		Cents:     cents,
		Timestamp: timestamp,
	}
}

func (c *Charge) Save(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO charge (account_id, cents, timestamp) VALUES ( ?, ?, ? )")
	if err != nil {
		fmt.Printf("error preparing statement: %v\n", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(c.AccountID, c.Cents, c.Timestamp)
	if err != nil {
		fmt.Printf("error creating %v: %v\n", c, err)
		return err
	}
	c.ID, err = res.LastInsertId()
	if err != nil {
		fmt.Printf("error creating %v: %v\n", c, err)
		return err
	}

	return nil
}

func GetBy(db *sql.DB, k string, v string) (Charge, error) {
	var c Charge

	stmt, err := db.Prepare(fmt.Sprintf("SELECT id, account_id, cents, timestamp FROM charge WHERE %v = ?", k))
	if err != nil {
		fmt.Printf("error preparing statement: %v\n", err)
		return c, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(v).Scan(&c.ID, &c.AccountID, &c.Cents, &c.Timestamp)
	if err == sql.ErrNoRows {
		return c, scroll.NotFoundError{
			Description: "charge not found",
		}
	}
	if err != nil {
		fmt.Printf("error reading charge(%v=%v): %v\n", k, v, err)
		return c, err
	}

	return c, nil
}

func Get(db *sql.DB, id int64) (Charge, error) {
	return GetBy(db, "id", strconv.FormatInt(id, 10))
}

func List(db *sql.DB, last int64, limit int64) ([]Charge, error) {
	var cl = make([]Charge, 0, limit)

	stmt, err := db.Prepare("SELECT id, account_id, cents, timestamp FROM charge WHERE id > ? LIMIT ?")
	if err != nil {
		fmt.Printf("error preparing statement: %v\n", err)
		return cl, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(last, limit)
	if err != nil {
		fmt.Printf("error listing charges(%v, %v)", last, limit)
		return cl, err
	}

	for rows.Next() {
		var c Charge
		err = rows.Scan(&c.ID, &c.AccountID, &c.Cents, &c.Timestamp)
		if err != nil {
			fmt.Printf("error listing charges(%v, %v)", last, limit)
			return cl, err
		}
		cl = append(cl, c)
	}

	return cl, nil
}
