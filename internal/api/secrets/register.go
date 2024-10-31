package secrets

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gitlab.com/xrs-cloud/lockbox/core/internal/secrets"
)

// SecretsService is the service layer that handles business logic for secrets.
// This package variable allows handlers to interact with the secret management service.
var SecretsService secrets.Service

// validate is a JSON validator to check JSON request bodies
var validate = validator.New()

// RegisterSecretsRoutes registers the HTTP routes for managing secrets.
// This function sets up the routes for creating, retrieving, updating, and deleting secrets.
// The routes are bound to handler functions that interact with the secrets service layer.
//
// Parameters:
// - router: The main router to which the secrets subrouter will be attached.
// - secretsService: The secrets service that will be used to handle the business logic related to secret management.
//
// Routes:
// - POST /secrets: Creates a new secret.
// - GET /secrets/{query}: Retrieves a secret by its UUID or key.
// - PUT /secrets/{query}: Updates an existing secret by its UUID or key.
// - DELETE /secrets/{query}: Deletes a secret by its UUID or key.
func RegisterSecretsRoutes(router *mux.Router, secretsService secrets.Service) {
	// Assign the provided secrets service to the package-level variable for use in the handler functions.
	SecretsService = secretsService

	// Create a subrouter for secret management under the /secrets path.
	secretsRouter := router.PathPrefix("/secrets").Subrouter()

	// Define the HTTP routes for managing secrets, and bind each route to its corresponding handler function.

	// POST /secrets: This route is used to create a new secret.
	secretsRouter.HandleFunc("", CreateSecret).Methods("POST")

	// GET /secrets/{query}: This route retrieves a secret by its UUID or unique key.
	secretsRouter.HandleFunc("/{query}", GetSecretByQuery).Methods("GET")

	// PUT /secrets/{query}: This route updates an existing secret.
	secretsRouter.HandleFunc("/{query}", UpdateSecret).Methods("PUT")

	// DELETE /secrets/{query}: This route deletes a secret by its UUID or key.
	secretsRouter.HandleFunc("/{query}", DeleteSecret).Methods("DELETE")
}
