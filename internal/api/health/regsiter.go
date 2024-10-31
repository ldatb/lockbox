package health

import "github.com/gorilla/mux"

// RegisterHealthRoutes registers routes for health checks.
// These endpoints are typically used to verify the application's status (readiness/liveness) in monitoring systems like Kubernetes.
func RegisterHealthRoutes(router *mux.Router) {
	// Create a subrouter for health check-related endpoints
	healthRouter := router.PathPrefix("/healthz").Subrouter()

	// Register the health check handler
	// The handler will respond to requests at /healthz, typically used for application health checks
	healthRouter.HandleFunc("", basicHealthCheck).Methods("GET")
	healthRouter.HandleFunc("/detailed", detailedHealthCheck).Methods("GET")
}
