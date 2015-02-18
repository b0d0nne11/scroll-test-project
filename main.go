package main

import (
	"net/http"

	"github.com/b0d0nne11/scroll-test-project/handlers"
	"github.com/mailgun/scroll"
	"github.com/mailgun/scroll/registry"
)

func main() {
	// Create the app
	appConfig := scroll.AppConfig{
		Name:       "scroll-test-project",
		ListenIP:   "0.0.0.0",
		ListenPort: 8080,
		Registry:   &registry.NopRegistry{},
	}
	app := scroll.NewAppWithConfig(appConfig)

	// Index
	app.AddHandler(scroll.Spec{
		Methods:    []string{"GET"},
		Paths:      []string{"/index"},
		RawHandler: http.HandlerFunc(handlers.ReplyNotImplemented),
	})

	// List accounts
	app.AddHandler(scroll.Spec{
		Methods:    []string{"GET"},
		Paths:      []string{"/api/v1/accounts/"},
		RawHandler: http.HandlerFunc(handlers.ReplyNotImplemented),
	})
	// Get account
	app.AddHandler(scroll.Spec{
		Methods:    []string{"GET"},
		Paths:      []string{"/api/v1/accounts/{accountId}"},
		RawHandler: http.HandlerFunc(handlers.ReplyNotImplemented),
	})

	// List charges
	app.AddHandler(scroll.Spec{
		Methods:    []string{"GET"},
		Paths:      []string{"/api/v1/chages/"},
		RawHandler: http.HandlerFunc(handlers.ReplyNotImplemented),
	})
	// Get charge
	app.AddHandler(scroll.Spec{
		Methods:    []string{"GET"},
		Paths:      []string{"/api/v1/charges/{chargeId}"},
		RawHandler: http.HandlerFunc(handlers.ReplyNotImplemented),
	})
	// Create charge
	app.AddHandler(scroll.Spec{
		Methods:    []string{"POST"},
		Paths:      []string{"/api/v1/charges/"},
		RawHandler: http.HandlerFunc(handlers.ReplyNotImplemented),
	})

	// Start the app
	app.Run()
}
