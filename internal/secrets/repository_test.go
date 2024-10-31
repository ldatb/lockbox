package secrets

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gitlab.com/xrs-cloud/lockbox/core/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Set up the database connection and return the repository instance
func setupTestRepository(t *testing.T) Repository {
	// Get database variables from env
	dbHost := utils.GetEnvOrFallback("DB_HOST", "localhost")
	dbPort := utils.GetEnvOrFallback("DB_PORT", "5432")
	dbUser := utils.GetEnvOrFallback("POSTGRES_USER", "testuser")
	dbPass := utils.GetEnvOrFallback("POSTGRES_PASSWORD", "testpassword")
	dbDatabase := utils.GetEnvOrFallback("POSTGRES_DB", "testdb")

	// Connect to database
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost,
		dbUser,
		dbPass,
		dbDatabase,
		dbPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	assert.NoError(t, err)

	// Make migration
	db.AutoMigrate(&Secret{})

	// Return repository
	return NewRepository(db)
}

// TestRepoSaveSecret tests saving a new secret in the database.
func TestRepoSaveSecret(t *testing.T) {
	repo := setupTestRepository(t)

	// Create a new secret
	secret := &Secret{
		ID:             uuid.New(),
		Key:            "key_TestRepoSaveSecret",
		EncryptedValue: "test_encrypted_value",
	}

	// Save the secret
	err := repo.Save(secret)
	assert.NoError(t, err)

	// Clean up
	repo.Delete(secret.ID)
}

// TestRepoNegativeSaveSecretDuplicatedKey tests saving a secret
// with a duplicated key
func TestRepoNegativeSaveSecretDuplicatedKey(t *testing.T) {
	repo := setupTestRepository(t)

	// Create a new secret
	secret := &Secret{
		ID:             uuid.New(),
		Key:            "key_TestRepoSaveSecret",
		EncryptedValue: "test_encrypted_value",
	}

	// Save the secret
	err := repo.Save(secret)
	assert.NoError(t, err)

	// Create the same secret again
	err = repo.Save(secret)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "duplicate key value violates unique constraint")

	// Clean up
	repo.Delete(secret.ID)
}

// TestRepoGetByID tests retrieving a secret by its UUID.
func TestRepoGetByID(t *testing.T) {
	repo := setupTestRepository(t)

	// Create and save a new secret
	secret := &Secret{
		ID:             uuid.New(),
		Key:            "test_TestRepoGetByID",
		EncryptedValue: "test_encrypted_value",
	}
	err := repo.Save(secret)
	assert.NoError(t, err)

	// Retrieve the secret by its UUID
	retrievedSecret, err := repo.GetByID(secret.ID)
	assert.NoError(t, err)
	assert.NotNil(t, retrievedSecret)
	assert.Equal(t, secret.ID, retrievedSecret.ID)
	assert.Equal(t, secret.Key, retrievedSecret.Key)

	// Clean up
	repo.Delete(secret.ID)
}

// TestRepoGetByKey tests retrieving a secret by its key.
func TestRepoGetByKey(t *testing.T) {
	repo := setupTestRepository(t)

	// Create and save a new secret
	secret := &Secret{
		ID:             uuid.New(),
		Key:            "test_TestRepoGetByKey",
		EncryptedValue: "test_encrypted_value",
	}
	err := repo.Save(secret)
	assert.NoError(t, err)

	// Retrieve the secret by its key
	retrievedSecret, err := repo.GetByKey(secret.Key)
	assert.NoError(t, err)
	assert.NotNil(t, retrievedSecret)
	assert.Equal(t, secret.Key, retrievedSecret.Key)

	// Clean up
	repo.Delete(secret.ID)
}

// TestRepoUpdateSecret tests updating an existing secret.
func TestRepoUpdateSecret(t *testing.T) {
	repo := setupTestRepository(t)

	// Create and save a new secret
	secret := &Secret{
		ID:             uuid.New(),
		Key:            "test_TestRepoUpdateSecret",
		EncryptedValue: "test_encrypted_value",
	}
	err := repo.Save(secret)
	assert.NoError(t, err)

	// Update the secret's encrypted value
	newEncryptedValue := "updated_encrypted_value"
	err = repo.Update(secret.ID, newEncryptedValue)
	assert.NoError(t, err)

	// Retrieve and check the updated value
	updatedSecret, err := repo.GetByID(secret.ID)
	assert.NoError(t, err)
	assert.Equal(t, newEncryptedValue, updatedSecret.EncryptedValue)

	// Clean up
	repo.Delete(secret.ID)
}

// TestRepoDeleteSecret tests deleting a secret from the database.
func TestRepoDeleteSecret(t *testing.T) {
	repo := setupTestRepository(t)

	// Create and save a new secret
	secret := &Secret{
		ID:             uuid.New(),
		Key:            "test_TestRepoDeleteSecret",
		EncryptedValue: "test_encrypted_value",
	}
	err := repo.Save(secret)
	assert.NoError(t, err)

	// Delete the secret
	err = repo.Delete(secret.ID)
	assert.NoError(t, err)

	// Attempt to retrieve the deleted secret (should return an error)
	_, err = repo.GetByID(secret.ID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}
