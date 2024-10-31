package secrets

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.com/xrs-cloud/lockbox/core/internal/global"
)

// MockLogger is a mock implementation of the Logger interface.
type MockLogger struct {
	mock.Mock
}

// Set up the repository and return the service
func setupTestService(t *testing.T) Service {
	repo := setupTestRepository(t)
	return NewService(repo)
}

// TestServiceCreateSecret tests the CreateSecret method
func TestServiceCreateSecret(t *testing.T) {
	service := setupTestService(t)
	global.Logger = logrus.New()

	// Create the Secret object
	id, key, err := service.CreateSecret(testKey, testPlainTextSecret, testMasterKey)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, testKey, key)

	// Cleanup
	service.DeleteSecret(id)
}

// TestServiceGetEncryptedSecretByID tests GetEncryptedSecretByID
func TestServiceGetEncryptedSecretByID(t *testing.T) {
	service := setupTestService(t)

	// Create the Secret object
	id, key, err := service.CreateSecret(testKey, testPlainTextSecret, testMasterKey)
	assert.NoError(t, err)

	// Get the secret by the ID
	retrievedSecret, err := service.GetEncryptedSecretByID(id)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, retrievedSecret)
	assert.Equal(t, id, retrievedSecret.ID.String())
	assert.Equal(t, key, retrievedSecret.Key)

	// Cleanup
	service.DeleteSecret(id)
}

// TestServiceGetEncryptedSecretByKey tests GetEncryptedSecretByID
func TestServiceGetEncryptedSecretByKey(t *testing.T) {
	service := setupTestService(t)

	// Create the Secret object
	id, key, err := service.CreateSecret(testKey, testPlainTextSecret, testMasterKey)
	assert.NoError(t, err)

	// Get the secret by the ID
	retrievedSecret, err := service.GetEncryptedSecretByKey(key)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, retrievedSecret)
	assert.Equal(t, id, retrievedSecret.ID.String())
	assert.Equal(t, key, retrievedSecret.Key)

	// Cleanup
	service.DeleteSecret(id)
}

// TestServiceDecryptSecret tests the DecryptSecret method
func TestServiceDecryptSecret(t *testing.T) {
	service := setupTestService(t)

	// Create the Secret object
	id, key, err := service.CreateSecret(testKey, testPlainTextSecret, testMasterKey)
	assert.NoError(t, err)

	// Get the secret by the ID
	retrievedSecret, err := service.GetEncryptedSecretByKey(key)
	assert.NoError(t, err)

	// Decrypt
	decryptedValue, err := service.DecryptSecret(*retrievedSecret, testMasterKey)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, testPlainTextSecret, decryptedValue)

	// Cleanup
	service.DeleteSecret(id)
}

// TestServiceNegativeDecryptSecret tests the DecryptSecret method with a wrong master key
func TestServiceNegativeDecryptSecret(t *testing.T) {
	service := setupTestService(t)
	global.Logger = logrus.New()

	// Create the Secret object
	id, key, err := service.CreateSecret(testKey, testPlainTextSecret, testMasterKey)
	assert.NoError(t, err)

	// Get the secret by the ID
	retrievedSecret, err := service.GetEncryptedSecretByKey(key)
	assert.NoError(t, err)

	// Decrypt
	_, err = service.DecryptSecret(*retrievedSecret, "not-the-true-key")

	// Assert
	assert.Error(t, err)

	// Cleanup
	service.DeleteSecret(id)
}

// TestServiceUpdateSecret tests the successful update of a secret.
func TestServiceUpdateSecret(t *testing.T) {
	service := setupTestService(t)

	// Create the Secret object
	id, _, err := service.CreateSecret(testKey, testPlainTextSecret, testMasterKey)
	assert.NoError(t, err)

	// Try updating the secret
	newPlainSecret := "this-secret-was-updated"
	err = service.UpdateSecret(id, newPlainSecret, testMasterKey)
	assert.NoError(t, err)

	// Get secret and make sure it changed
	retrievedSecret, err := service.GetEncryptedSecretByID(id)
	assert.NoError(t, err)
	decryptedValue, err := service.DecryptSecret(*retrievedSecret, testMasterKey)
	assert.NoError(t, err)

	// Assert secrets match
	assert.Equal(t, newPlainSecret, decryptedValue)

	// Cleanup
	service.DeleteSecret(id)
}

// TestServiceDeleteSecret tests the successful deletion of a secret
func TestServiceDeleteSecret(t *testing.T) {
	service := setupTestService(t)

	// Create the Secret object
	id, _, err := service.CreateSecret(testKey, testPlainTextSecret, testMasterKey)
	assert.NoError(t, err)

	// Use the delete method
	err = service.DeleteSecret(id)

	// Assert
	assert.NoError(t, err)
}
