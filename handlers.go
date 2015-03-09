package main

import (
	"math"
	"net/http"

	"github.com/b0d0nne11/scroll-test-project/models/account"
	"github.com/b0d0nne11/scroll-test-project/models/charge"
	"github.com/mailgun/scroll"
)

func createCharge(w http.ResponseWriter, r *http.Request, params map[string]string) (interface{}, error) {
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
	}
	if err != nil {
		return nil, err
	}

	return charge.New(a, cents, timestamp).Save()
}

func getCharge(w http.ResponseWriter, r *http.Request, params map[string]string) (interface{}, error) {
	return charge.Get(params["id"])
}

func listCharges(w http.ResponseWriter, r *http.Request, params map[string]string) (interface{}, error) {
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

func getAccount(w http.ResponseWriter, r *http.Request, params map[string]string) (interface{}, error) {
	return account.Get(params["id"])
}

func listAccounts(w http.ResponseWriter, r *http.Request, params map[string]string) (interface{}, error) {
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
