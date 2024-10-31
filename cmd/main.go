package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"gitlab.com/xrs-cloud/lockbox/core/internal/api"
	"gitlab.com/xrs-cloud/lockbox/core/internal/config"
	"gitlab.com/xrs-cloud/lockbox/core/internal/database"
	"gitlab.com/xrs-cloud/lockbox/core/internal/global"
	app_log "gitlab.com/xrs-cloud/lockbox/core/internal/logger"
)

func main() {
	// Define command-line flags for configuration file path
	configFile := flag.String("config-file", "/etc/lockbox/lockbox.conf", "Path to the configuration file (.conf)")
	flag.Parse()

	// Load application configuration
	// This step initializes all necessary settings for the application (server, database, logging, etc.)
	config, err := config.LoadConfig(*configFile)
	if err != nil {
		// Use log.Fatalf to immediately exit if configuration loading fails
		log.Fatalf("Error loading configuration file: %v", err)
	}

	// Initialize and configure the application logger
	// The logger is globally accessible through the `global` package
	logger := app_log.InitLogger(config.Logging)
	global.Logger = logger

	// Establish a connection to the database
	// If the connection cannot be established, the application will log and exit
	logger.Infof("Creating connection to the database at %s", config.Database.Host)
	db := database.InitDatabase(config.Database)
	global.Database = db

	// Setup the API router for handling HTTP requests
	router := api.SetupRouter()

	// Start the HTTP server
	// The server runs on the address and port specified in the configuration file
	// If the server fails to start, the application logs the error and exits
	serverAddressAndPort := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)
	logger.Infof("Starting server on %s", serverAddressAndPort)
	if err := http.ListenAndServe(serverAddressAndPort, router); err != nil {
		// Use logger.Fatalf to log the error and terminate the application
		logger.Fatalf("Error starting server: %v", err)
	}
}
