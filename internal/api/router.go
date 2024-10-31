package api

import (
	"github.com/gorilla/mux"
	health_handler "gitlab.com/xrs-cloud/lockbox/core/internal/api/health"
	"gitlab.com/xrs-cloud/lockbox/core/internal/api/middleware"
	secrets_handler "gitlab.com/xrs-cloud/lockbox/core/internal/api/secrets"
	"gitlab.com/xrs-cloud/lockbox/core/internal/global"
	"gitlab.com/xrs-cloud/lockbox/core/internal/secrets"
)

// SetupRouter initializes the router and defines the routes for all services.
// This function sets up the base router, applies any global middleware
// and registers all service-specific routes.
func SetupRouter() *mux.Router {
	// Initialize a new router using Gorilla Mux
	router := mux.NewRouter()

	// Apply global middleware for security, logging, CORS, etc.
	global.Logger.Info("Adding middlewares to router")
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.AuthenticationMiddleware)
	router.Use(middleware.CORSMiddleware)

	// Initialize the repositories
	global.Logger.Info("Initializing repositories")
	secretsRepository := secrets.NewRepository(global.Database)

	// Initialize the services
	global.Logger.Info("Initializing services")
	secretsService := secrets.NewService(secretsRepository)

	// Register service-specific routes
	// Each group of routes is handled by a dedicated function to maintain separation of concerns
	global.Logger.Info("Registering routes")
	health_handler.RegisterHealthRoutes(router)
	secrets_handler.RegisterSecretsRoutes(router, secretsService)

	// Return the configured router
	return router
}
