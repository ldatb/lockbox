package secrets

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Repository interface defines methods for database interactions related to secrets.
// It encapsulates basic CRUD operations for managing secrets in the database.
type Repository interface {
	// Saves a new secret to the database
	Save(secret *Secret) error

	// Retrieves a secret by its UUID
	GetByID(secretID uuid.UUID) (*Secret, error)

	// Retrieves a secret by its key
	GetByKey(key string) (*Secret, error)

	// Updates the encrypted value of a secret
	Update(secretID uuid.UUID, newValue string) error

	// Deletes a secret from the database by its UUID
	Delete(secretID uuid.UUID) error
}

type repository struct {
	db *gorm.DB // The database connection, injected into the repository
}

// NewRepository creates a new instance of the secrets repository.
// The repository is initialized with a GORM database connection.
func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

// Save inserts a new secret record into the database.
// It takes a Secret object that contains the ID, encrypted value, and timestamps.
//
// Parameters:
// - secret: The Secret model containing the encrypted data to be stored.
//
// Returns:
// - error: Returns an error if the insertion fails, otherwise nil.
func (r *repository) Save(secret *Secret) error {
	return r.db.Create(secret).Error
}

// GetByID retrieves a secret from the database by its UUID.
// It searches for a secret with the given UUID and returns it if found.
//
// Parameters:
// - secretID: The UUID of the secret to retrieve.
//
// Returns:
// - Secret: The retrieved Secret model.
// - error: Returns an error if no secret with the given ID is found or if the query fails.
func (r *repository) GetByID(secretID uuid.UUID) (*Secret, error) {
	var secret *Secret
	err := r.db.First(&secret, "id = ?", secretID).Error
	return secret, err
}

// GetByKey retrieves a secret from the database by its key.
// It searches for a secret with the given key and returns it if found.
//
// Parameters:
// - key: The key of the secret to retrieve.
//
// Returns:
// - Secret: The retrieved Secret model.
// - error: Returns an error if no secret with the given ID is found or if the query fails.
func (r *repository) GetByKey(key string) (*Secret, error) {
	var secret *Secret
	err := r.db.First(&secret, "key = ?", key).Error
	return secret, err
}

// Update modifies the encrypted value of an existing secret.
// It looks up the secret by its UUID and updates its EncryptedValue field.
//
// Parameters:
// - secretID: The UUID of the secret to update.
// - newValue: The new encrypted value that will replace the old one.
//
// Returns:
// - error: Returns an error if the update fails, otherwise nil.
func (r *repository) Update(secretID uuid.UUID, newValue string) error {
	return r.db.Model(&Secret{}).Where("id = ?", secretID).Update("encrypted_value", newValue).Error
}

// Delete removes a secret from the database by its UUID.
// It performs a hard delete of the record identified by the given UUID.
//
// Parameters:
// - secretID: The UUID of the secret to delete.
//
// Returns:
// - error: Returns an error if the deletion fails, otherwise nil.
func (r *repository) Delete(secretID uuid.UUID) error {
	return r.db.Delete(&Secret{}, "id = ?", secretID).Error
}
