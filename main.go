package main

import (
	"flag"
	"fmt"

	"github.com/b0d0nne11/scroll-test-project/db"
	"github.com/mailgun/scroll"
	"github.com/mailgun/scroll/registry"
)

var host *string = flag.String("h", "0.0.0.0", "Address to listen on")
var port *int = flag.Int("p", 8080, "Port to listen on")

func main() {
	// Parse command line options
	flag.Parse()

	// Create the app
	appConfig := scroll.AppConfig{
		Name:       "scroll-test-project",
		ListenIP:   *host,
		ListenPort: *port,
		Registry:   &registry.NopRegistry{},
	}
	app := scroll.NewAppWithConfig(appConfig)

	// Create the db pool
	dbh, err := db.Get()
	if err != nil {
		panic(fmt.Sprintf("error connecting to db: %v\n", err))
	}
	defer dbh.Close()

	// List accounts
	app.AddHandler(scroll.Spec{
		Methods: []string{"GET"},
		Paths:   []string{"/api/v1/accounts/"},
		Handler: ListAccounts,
	})
	// Get account
	app.AddHandler(scroll.Spec{
		Methods: []string{"GET"},
		Paths:   []string{"/api/v1/accounts/{id}"},
		Handler: GetAccount,
	})

	// List charges
	app.AddHandler(scroll.Spec{
		Methods: []string{"GET"},
		Paths:   []string{"/api/v1/charges/"},
		Handler: ListCharges,
	})
	// Get charge
	app.AddHandler(scroll.Spec{
		Methods: []string{"GET"},
		Paths:   []string{"/api/v1/charges/{id}"},
		Handler: GetCharge,
	})
	// Create charge
	app.AddHandler(scroll.Spec{
		Methods: []string{"POST"},
		Paths:   []string{"/api/v1/charges/"},
		Handler: CreateCharge,
	})

	// Start the app
	app.Run()
}
