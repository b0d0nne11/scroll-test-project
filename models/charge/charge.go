package charge

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/mailgun/scroll"
)

const Schema string = `
CREATE TABLE IF NOT EXISTS charge (
	id integer NOT NULL AUTO_INCREMENT,
  cents int,
	timestamp datetime,
	PRIMARY KEY (id)
);
`

type Charge struct {
	ID        int64
	Cents     uint64
	Timestamp time.Time
}

func New(cents uint64, timestamp time.Time) Charge {
	return Charge{
		Cents:     cents,
		Timestamp: timestamp,
	}
}

func (c Charge) Save(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO charge (cents, timestamp) VALUES ( ?, ? )")
	if err != nil {
		fmt.Printf("error preparing statement: %v\n", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(c.Cents, c.Timestamp)
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

func Get(db *sql.DB, id int64) (Charge, error) {
	var c Charge

	stmt, err := db.Prepare("SELECT id, cents, timestamp FROM charge WHERE id = ?")
	if err != nil {
		fmt.Printf("error preparing statement: %v\n", err)
		return c, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&c.ID, &c.Cents, &c.Timestamp)
	if err == sql.ErrNoRows {
		return c, scroll.NotFoundError{
			Description: "charge not found",
		}
	}
	if err != nil {
		fmt.Printf("error reading charge(%v): %v\n", id, err)
		return c, err
	}

	return c, nil
}
