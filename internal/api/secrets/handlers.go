package secrets

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gitlab.com/xrs-cloud/lockbox/core/internal/secrets"
	"gitlab.com/xrs-cloud/lockbox/core/internal/utils"
)

// CreateSecret handles creating a new secret and storing it in the database.
// It expects a JSON body with the "key" and "plain_text_secret" fields, and the secret is encrypted using
// the "MASTER_CRYPTO_PASS" environment variable.
// If successful, it returns the key of the newly created secret.
//
// Expected JSON request body:
//
//	{
//	    "secret_key": "unique_key_for_secret",
//	    "secret_value": "sensitive_value_to_store"
//	}
//
// Responses:
// - 201 Created: Returns the key of the newly created secret.
// - 400 Bad Request: Returns if the request body is invalid.
// - 500 Internal Server Error: Returns if the secret creation fails.
func CreateSecret(w http.ResponseWriter, r *http.Request) {
	// Get JSON request body
	var req struct {
		SecretKey   string `json:"secret_key" validate:"required"`
		SecretValue string `json:"secret_value" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	// Validate the decoded struct using the validator package
	if err := validate.Struct(&req); err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	// Get master key from the environment
	masterCryptoPass := os.Getenv("MASTER_CRYPTO_PASS")

	// Create the secret using the service layer
	secretID, secretKey, err := SecretsService.CreateSecret(req.SecretKey, req.SecretValue, masterCryptoPass)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create secret"})
		return
	}

	// Create and return presenter
	presenter := &SecretResponseUUID{
		ID:  secretID,
		Key: secretKey,
	}
	utils.WriteJSONResponse(w, http.StatusCreated, presenter)
}

// GetSecretByQuery retrieves an encrypted secret based on the provided query (UUID or key).
// The secret is decrypted using the "MASTER_CRYPTO_PASS" environment variable before being returned.
// It supports lookup by either the UUID or a unique key, depending on the query value.
//
// Responses:
// - 200 OK: Returns the decrypted secret.
// - 400 Bad Request: Returns if the query is missing from the URL.
// - 404 Not Found: Returns if the secret cannot be found using the given query.
// - 500 Internal Server Error: Returns if decryption or retrieval fails.
func GetSecretByQuery(w http.ResponseWriter, r *http.Request) {
	// Get query from URL
	query := mux.Vars(r)["query"]
	if query == "" {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Missing query in request URL"})
		return
	}

	// Get secret based on query
	secret, err := getSecretFromQuery(query)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusNotFound, map[string]string{"error": "Secret not found"})
		return
	}

	// Get master key from the environment
	masterCryptoPass := os.Getenv("MASTER_CRYPTO_PASS")

	// Decrypt the secret
	decryptedSecret, err := SecretsService.DecryptSecret(*secret, masterCryptoPass)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Something went wrong"})
		return
	}

	// Create and return presenter
	presenter := &SecretResponsePlain{
		Key:   secret.Key,
		Value: decryptedSecret,
	}
	utils.WriteJSONResponse(w, http.StatusOK, presenter)
}

// UpdateSecret handles updating an existing secret based on the provided query (UUID or key).
// It expects a JSON body with the new "plain_text_secret" value, which will replace the old encrypted secret.
// The secret is re-encrypted with the "MASTER_CRYPTO_PASS" environment variable.
//
// Expected JSON request body:
//
//	{
//	    "secret_value": "new_secret_value"
//	}
//
// Responses:
// - 200 OK: Returns if the secret was successfully updated.
// - 400 Bad Request: Returns if the request body or query is invalid.
// - 404 Not Found: Returns if the secret cannot be found using the given query.
// - 500 Internal Server Error: Returns if the update operation fails.
func UpdateSecret(w http.ResponseWriter, r *http.Request) {
	// Get query from URL
	query := mux.Vars(r)["query"]
	if query == "" {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Missing query in request URL"})
		return
	}

	// Get JSON request body
	var req struct {
		NewSecretValue string `json:"secret_value" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	// Validate the decoded struct using the validator package
	if err := validate.Struct(&req); err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	// Get secret based on query
	secret, err := getSecretFromQuery(query)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusNotFound, map[string]string{"error": "Secret not found"})
		return
	}

	// Get master key from the environment
	masterCryptoPass := os.Getenv("MASTER_CRYPTO_PASS")

	// Update the secret with the new plain text secret
	if err := SecretsService.UpdateSecret(secret.ID.String(), req.NewSecretValue, masterCryptoPass); err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to update secret"})
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, map[string]string{"message": "Secret updated successfully"})
}

// DeleteSecret handles deleting an existing secret based on the provided query (UUID or key).
// It removes the secret from the database.
//
// Responses:
// - 200 OK: Returns if the secret was successfully deleted.
// - 400 Bad Request: Returns if the query is missing from the URL.
// - 404 Not Found: Returns if the secret cannot be found using the given query.
// - 500 Internal Server Error: Returns if the delete operation fails.
func DeleteSecret(w http.ResponseWriter, r *http.Request) {
	// Get query from URL
	query := mux.Vars(r)["query"]
	if query == "" {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Missing query in request URL"})
		return
	}

	// Get secret based on query
	secret, err := getSecretFromQuery(query)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusNotFound, map[string]string{"error": "Secret not found"})
		return
	}

	// Delete the secret
	if err := SecretsService.DeleteSecret(secret.ID.String()); err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to delete secret"})
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, map[string]string{"message": "Secret deleted successfully"})
}

// getSecretFromQuery determines whether the query is a UUID or a unique key and retrieves the corresponding secret.
// If the query is a valid UUID, it retrieves the secret by ID; otherwise, it retrieves the secret by key.
//
// Parameters:
// - query: The UUID or key used to look up the secret.
//
// Returns:
// - *secrets.Secret: The retrieved secret, or an error if not found or if the query is invalid.
func getSecretFromQuery(query string) (*secrets.Secret, error) {
	// Define if it's a UUID or a key
	queryType := "uuid"
	_, err := uuid.Parse(query)
	if err != nil {
		queryType = "key"
	}

	// Get secret based on UUID or key
	var secret *secrets.Secret
	switch queryType {
	case "uuid":
		secret, err = SecretsService.GetEncryptedSecretByID(query)
	case "key":
		secret, err = SecretsService.GetEncryptedSecretByKey(query)
	}

	return secret, err
}
