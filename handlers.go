package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/b0d0nne11/scroll-test-project/models/account"
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
	aName := r.FormValue("account_name")
	if aName == "" {
		return nil, scroll.InvalidFormatError{
			Field: "account_name",
			Value: r.FormValue("account_name"),
		}
	}

	a, err := account.GetBy("name", aName)
	switch err.(type) {
	case scroll.NotFoundError:
		a = account.New(aName)
		err = a.Save()
	}

	c := charge.New(a.ID, cents, timestamp)
	err = c.Save()

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

	c, err := charge.Get(id)

	return c, err
}

func ListCharges(w http.ResponseWriter, r *http.Request, params map[string]string) (interface{}, error) {
	last, err := strconv.ParseInt(r.FormValue("last"), 10, 64)
	if err != nil {
		last = 0
	}
	limit, err := strconv.ParseInt(r.FormValue("limit"), 10, 64)
	if err != nil {
		limit = 100
	}

	cl, err := charge.List(last, limit)

	return cl, err
}

func GetAccount(w http.ResponseWriter, r *http.Request, params map[string]string) (interface{}, error) {
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		return nil, scroll.InvalidFormatError{
			Field: "id",
			Value: params["id"],
		}
	}

	a, err := account.Get(id)

	return a, err
}

func ListAccounts(w http.ResponseWriter, r *http.Request, params map[string]string) (interface{}, error) {
	last, err := strconv.ParseInt(r.FormValue("last"), 10, 64)
	if err != nil {
		last = 0
	}
	limit, err := strconv.ParseInt(r.FormValue("limit"), 10, 64)
	if err != nil {
		limit = 100
	}

	al, err := account.List(last, limit)

	return al, err
}
