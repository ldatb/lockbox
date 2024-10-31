package secrets

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gitlab.com/xrs-cloud/lockbox/core/internal/global"
)

// Secret represents a model that stores sensitive information in encrypted form.
// The sensitive data is encrypted before being saved in the database, ensuring security
// of the stored information. The model also tracks creation and update timestamps for audit purposes.
type Secret struct {
	// ID is the unique identifier for each secret record.
	// This field is the primary key in the database.
	ID uuid.UUID `gorm:"primaryKey"`

	// Key is a unique identifier for the secret, used to identify and retrieve the secret without relying on the UUID.
	Key string `gorm:"unique"`

	// EncryptedValue holds the encrypted version of the sensitive data.
	// The actual secret (like an API key, password, or token) is encrypted using AES-256 encryption,
	// and the encrypted result is stored as a string in this field.
	EncryptedValue string `gorm:"not null"`

	// CreatedAt stores the timestamp of when the secret was created.
	// This field is automatically populated by GORM when a new record is inserted into the database.
	CreatedAt time.Time `gorm:"autoCreateTime"`

	// UpdatedAt stores the timestamp of when the secret was last updated.
	// This field is automatically updated by GORM every time the secret is modified.
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// CreateSecretModel encrypts the provided plain text secret and returns the model
//
// Parameters:
// - plainText: The sensitive data (e.g., API key, password) that needs to be encrypted and stored.
// - masterKey: The passphrase used to encrypt the secret.
//
// Returns:
// - The created Secret model.
// - An error if anything goes wrong during the encryption.
func CreateSecretModel(key, plainTextSecret, masterKey string) (*Secret, error) {
	// Encrypt the plainText using the provided masterKey
	encryptedValue, err := Encrypt(plainTextSecret, masterKey)
	if err != nil {
		err = fmt.Errorf("failed to encrypt secret: %v", err)
		global.Logger.Error(err)
		return nil, err
	}

	// Create a new Secret model with the encrypted value
	secret := &Secret{
		ID:             uuid.New(),     // Generate a new UUID for the secret
		Key:            key,            // Store the key as a plain text
		EncryptedValue: encryptedValue, // Store the encrypted secret
	}

	// Return the created secret model
	return secret, nil
}
