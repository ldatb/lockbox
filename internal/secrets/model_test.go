package secrets

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	testKey             = "my-secret-key"
	testPlainTextSecret = "super-secret"
	testMasterKey       = "test-master-key-1234"
)

// TestCreateSecretModelSuccess tests successful creation of the Secret model.
func TestCreateSecretModel(t *testing.T) {
	// Create Secrets model
	secret, err := CreateSecretModel(testKey, testPlainTextSecret, testMasterKey)
	assert.NoError(t, err)
	assert.NotNil(t, secret)

	// Validate fields
	assert.Equal(t, testKey, secret.Key)
	assert.NotEqual(t, "encrypted_super-secret", secret.EncryptedValue)
	assert.NotEqual(t, uuid.Nil, secret.ID)
	assert.True(t, time.Now().After(secret.CreatedAt))
	assert.True(t, time.Now().After(secret.UpdatedAt))
}
