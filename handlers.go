package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/b0d0nne11/scroll-test-project/models/charge"
	"github.com/mailgun/scroll"
)

func ReplyNotImplemented(w http.ResponseWriter, r *http.Request) {
	scroll.Reply(w, scroll.Response{"message": "Not Implemented"}, http.StatusNotImplemented)
}

func CreateCharge(w http.ResponseWriter, r *http.Request, params map[string]string) (interface{}, error) {
	cents, err := strconv.ParseUint(r.FormValue("cents"), 10, 64)
	if err != nil {
		return nil, scroll.InvalidFormatError{
			Field: "cents",
			Value: r.FormValue("cents"),
		}
	}
	timestamp, err := time.Parse("2006-01-02T15:04:05", r.FormValue("timestamp"))
	if err != nil {
		return nil, scroll.InvalidFormatError{
			Field: "timestamp",
			Value: r.FormValue("timestamp"),
		}
	}

	c := charge.New(cents, timestamp)
	err = c.Save(db)

	return c, err
}

func GetCharge(w http.ResponseWriter, r *http.Request, params map[string]string) (interface{}, error) {
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		return nil, scroll.InvalidFormatError{
			Field: "id",
			Value: params["id"],
		}
	}

	c, err := charge.Get(db, id)

	return c, err
}
