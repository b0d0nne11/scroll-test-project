package main

import (
	"flag"
	"fmt"

	"github.com/mailgun/scroll"
	"github.com/mailgun/scroll/registry"
)

var host *string = flag.String("h", "0.0.0.0", "Address to listen on")
var port *int = flag.Int("p", 8080, "Port to listen on")
var dbconn *string = flag.String("db", "username:password@address/database", "Database connection string")

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

	// Create a db pool
	db, err := GetDB()
	if err != nil {
		panic(fmt.Sprintf("error connecting to db: %v\n", err))
	}
	defer db.Close()

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
