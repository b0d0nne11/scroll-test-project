package main

import (
	"math"
	"net/http"
	"strconv"

	"github.com/b0d0nne11/scroll-test-project/models/account"
	"github.com/b0d0nne11/scroll-test-project/models/charge"
	"github.com/mailgun/scroll"
)

func ReplyNotImplemented(w http.ResponseWriter, r *http.Request) {
	scroll.Reply(w, scroll.Response{"message": "Not Implemented"}, http.StatusNotImplemented)
}

func CreateCharge(w http.ResponseWriter, r *http.Request, params map[string]string) (interface{}, error) {
	cents, err := scroll.GetIntField(r, "cents")
	if err != nil {
		return nil, err
	}
	timestamp, err := scroll.GetTimestampField(r, "timestamp")
	if err != nil {
		return nil, err
	}
	aName, err := scroll.GetStringField(r, "account_name")
	if err != nil {
		return nil, err
	}

	a, err := account.FindByName(aName)
	switch err.(type) {
	case scroll.NotFoundError:
		a, err = account.New(aName).Save()
		if err != nil {
			return nil, err
		}
	}

	return charge.New(a.ID, cents, timestamp).Save()
}

func GetCharge(w http.ResponseWriter, r *http.Request, params map[string]string) (interface{}, error) {
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		return nil, scroll.InvalidFormatError{
			Field: "id",
			Value: params["id"],
		}
	}

	return charge.Get(id)
}

func ListCharges(w http.ResponseWriter, r *http.Request, params map[string]string) (interface{}, error) {
	last, err := scroll.GetIntField(r, "last")
	if err != nil {
		last = 0
	}
	limit, err := scroll.GetIntField(r, "limit")
	if err != nil {
		limit = 100
	}
	limit = int(math.Min(1000, math.Max(0, float64(limit))))

	return charge.List(last, limit)
}

func GetAccount(w http.ResponseWriter, r *http.Request, params map[string]string) (interface{}, error) {
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		return nil, scroll.InvalidFormatError{
			Field: "id",
			Value: params["id"],
		}
	}

	return account.Get(id)
}

func ListAccounts(w http.ResponseWriter, r *http.Request, params map[string]string) (interface{}, error) {
	last, err := scroll.GetIntField(r, "last")
	if err != nil {
		last = 0
	}
	limit, err := scroll.GetIntField(r, "limit")
	if err != nil {
		limit = 100
	}
	limit = int(math.Min(1000, math.Max(0, float64(limit))))

	return account.List(last, limit)
}
