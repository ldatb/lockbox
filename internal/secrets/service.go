package secrets

import (
	"fmt"

	"github.com/google/uuid"
	"gitlab.com/xrs-cloud/lockbox/core/internal/global"
)

// Service interface defines the business logic for handling secrets.
type Service interface {
	// CreateSecret encrypts the plainTextSecret using the provided masterKey and stores it in the database.
	// The secret is identified by a unique key for easy retrieval.
	// Returns the key or an error if something goes wrong.
	CreateSecret(key, plainTextSecret, masterKey string) (string, string, error)

	// GetEncryptedSecretByID retrieves an encrypted secret from the database using its UUID.
	// Decryption is deferred until the caller specifically requests it.
	// Returns the Secret model or an error if something goes wrong.
	GetEncryptedSecretByID(secretID string) (*Secret, error)

	// GetEncryptedSecretByKey retrieves an encrypted secret from the database using its unique Key.
	// Decryption is deferred until the caller specifically requests it.
	// Returns the Secret model or an error if something goes wrong.
	GetEncryptedSecretByKey(key string) (*Secret, error)

	// DecryptSecret decrypts the EncryptedValue of the Secret using the provided masterKey.
	// Returns the decrypted secret or an error if decryption fails.
	DecryptSecret(secret Secret, masterKey string) (string, error)

	// UpdateSecret updates the encrypted value of an existing secret using its unique key.
	// It re-encrypts the provided plainTextSecret and stores the new value in the database.
	// Returns an error if the update fails.
	UpdateSecret(secretID, plainTextSecret, masterKey string) error

	// DeleteSecret deletes a secret from the database by its UUID.
	// Returns an error if deletion fails.
	DeleteSecret(secretID string) error
}

type service struct {
	repo Repository
}

// NewService creates a new secret service.
func NewService(repo Repository) Service {
	return &service{repo}
}

// CreateSecret encrypts a secret and stores it in the database.
// This function takes a key (used to identify the secret), the plain-text secret, and the masterKey for encryption.
// Returns the key of the created secret or an error if something goes wrong.
func (s *service) CreateSecret(key, plainTextSecret, masterKey string) (string, string, error) {
	// Create the Secret model
	secret, err := CreateSecretModel(key, plainTextSecret, masterKey)
	if err != nil {
		err = fmt.Errorf("failed to create secret: %v", err)
		global.Logger.Error(err)
		return "", "", err
	}

	// Save the secret in the repository
	if err := s.repo.Save(secret); err != nil {
		err = fmt.Errorf("failed to store secret in the database: %v", err)
		global.Logger.Error(err)
		return "", "", err
	}

	return secret.ID.String(), secret.Key, nil
}

// GetEncryptedSecretByID retrieves an encrypted secret from the database using its UUID.
func (s *service) GetEncryptedSecretByID(secretID string) (*Secret, error) {
	// Convert the string ID to a UUID
	parserSecretID, err := uuid.Parse(secretID)
	if err != nil {
		err = fmt.Errorf("invalid UUID format: %v", err)
		global.Logger.Error(err)
		return &Secret{}, err
	}

	// Retrieve the encrypted secret from the repository
	secret, err := s.repo.GetByID(parserSecretID)
	if err != nil {
		err = fmt.Errorf("failed to retrieve secret by ID '%s': %v", parserSecretID, err)
		global.Logger.Debug(err)
		return &Secret{}, err
	}

	return secret, nil
}

// GetEncryptedSecretByKey retrieves an encrypted secret from the database using its unique Key.
func (s *service) GetEncryptedSecretByKey(key string) (*Secret, error) {
	// Retrieve the encrypted secret from the repository using its unique key
	secret, err := s.repo.GetByKey(key)
	if err != nil {
		err = fmt.Errorf("failed to retrieve secret by key '%s': %v", key, err)
		global.Logger.Debug(err)
		return &Secret{}, err
	}

	return secret, nil
}

// DecryptSecret decrypts the EncryptedValue of the Secret using the provided masterKey.
func (s *service) DecryptSecret(secret Secret, masterKey string) (string, error) {
	// Decrypt the secret using the masterKey
	decryptedValue, err := Decrypt(secret.EncryptedValue, masterKey)
	if err != nil {
		err = fmt.Errorf("failed to decrypt secret: %v", err)
		global.Logger.Error(err)
		return "", err
	}

	return decryptedValue, nil
}

// UpdateSecret updates the encrypted value of an existing secret.
// It re-encrypts the provided plainTextSecret and updates the secret in the database using the unique key.
func (s *service) UpdateSecret(secretID, plainTextSecret, masterKey string) error {
	// Convert the string ID to a UUID
	parserSecretID, err := uuid.Parse(secretID)
	if err != nil {
		err = fmt.Errorf("invalid UUID format: %v", err)
		global.Logger.Error(err)
		return err
	}

	// Encrypt the new plain-text secret
	encryptedValue, err := Encrypt(plainTextSecret, masterKey)
	if err != nil {
		err = fmt.Errorf("failed to encrypt secret: %v", err)
		global.Logger.Error(err)
		return err
	}

	// Update the secret in the repository using its unique key
	if err := s.repo.Update(parserSecretID, encryptedValue); err != nil {
		err = fmt.Errorf("failed to update secret: %v", err)
		global.Logger.Error(err)
		return err
	}

	return nil
}

// DeleteSecretByID deletes a secret from the database using its UUID.
func (s *service) DeleteSecret(secretID string) error {
	// Convert the string ID to a UUID
	parserSecretID, err := uuid.Parse(secretID)
	if err != nil {
		err = fmt.Errorf("invalid UUID format: %v", err)
		global.Logger.Error(err)
		return err
	}

	// Delete the secret from the repository using its UUID
	if err := s.repo.Delete(parserSecretID); err != nil {
		err = fmt.Errorf("failed to delete secret by ID: %v", err)
		global.Logger.Error(err)
		return err
	}

	return nil
}
