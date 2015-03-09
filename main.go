// Package main initializes and runs the application
package main

import (
	"flag"
	"fmt"

	"github.com/b0d0nne11/scroll-test-project/db"
	"github.com/mailgun/cfg"
	"github.com/mailgun/log"
	"github.com/mailgun/scroll"
	"github.com/mailgun/scroll/registry"
)

var host *string = flag.String("h", "0.0.0.0", "Address to listen on")
var port *int = flag.Int("p", 8080, "Port to listen on")
var confPath *string = flag.String("c", "./conf.yml", "Path to conf file")

// Config contains application configration parameters
type Config struct {
	Logging  []*log.LogConfig
	Database db.DatabaseConfig
}

func main() {
	// Parse command line options
	flag.Parse()

	// Parse configuration file
	conf := Config{}
	err := cfg.LoadConfig(*confPath, &conf)
	if err != nil {
		panic(fmt.Sprintf("error loading conf file: %v\n", err))
	}

	// Initialize the logging package
	log.Init(conf.Logging)
	log.SetSeverity(log.SeverityInfo) // TODO: make this configurable

	// Create the app
	appConfig := scroll.AppConfig{
		Name:       "scroll-test-project",
		ListenIP:   *host,
		ListenPort: *port,
		Registry:   &registry.NopRegistry{},
	}
	app := scroll.NewAppWithConfig(appConfig)

	// Initialize the db session
	dbSession, err := db.Init(conf.Database)
	if err != nil {
		panic(fmt.Sprintf("error connecting to db: %v\n", err))
	}
	defer dbSession.Close()

	// List accounts
	app.AddHandler(scroll.Spec{
		Methods: []string{"GET"},
		Paths:   []string{"/api/v1/accounts/"},
		Handler: listAccounts,
	})
	// Get account
	app.AddHandler(scroll.Spec{
		Methods: []string{"GET"},
		Paths:   []string{"/api/v1/accounts/{id}"},
		Handler: getAccount,
	})

	// List charges
	app.AddHandler(scroll.Spec{
		Methods: []string{"GET"},
		Paths:   []string{"/api/v1/charges/"},
		Handler: listCharges,
	})
	// Get charge
	app.AddHandler(scroll.Spec{
		Methods: []string{"GET"},
		Paths:   []string{"/api/v1/charges/{id}"},
		Handler: getCharge,
	})
	// Create charge
	app.AddHandler(scroll.Spec{
		Methods: []string{"POST"},
		Paths:   []string{"/api/v1/charges/"},
		Handler: createCharge,
	})

	// Start the app
	app.Run()
}
