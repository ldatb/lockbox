package secrets

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"gitlab.com/xrs-cloud/lockbox/core/internal/global"
	"io"
)

// Encrypt encrypts the provided plain-text secret using AES-256 GCM (Galois/Counter Mode) encryption.
// AES-256 GCM is a secure symmetric encryption algorithm that provides both confidentiality and integrity.
//
// Parameters:
// - plainText: The secret or sensitive data that needs to be encrypted (e.g., API key, password).
// - masterKey: A passphrase that is used to generate the encryption key via SHA-256.
//
// Returns:
// - A hex-encoded string of the encrypted data, which includes the encrypted secret and the nonce.
// - An error if encryption fails at any step.
//
// Encryption Process:
// 1. The masterKey is hashed using SHA-256 to generate a 32-byte key suitable for AES-256 encryption.
// 2. AES-256 GCM is initialized as the encryption mode.
// 3. A unique nonce (number used once) is generated for this encryption operation, ensuring uniqueness even when encrypting the same data multiple times.
// 4. The plainText is encrypted using AES-256 GCM, and the result (cipherText) is combined with the nonce.
// 5. The nonce and encrypted secret are returned as a hex-encoded string.
func Encrypt(plainText, masterKey string) (string, error) {
	// Create a new AES cipher block using the hashed master key
	block, err := aes.NewCipher(createHash(masterKey))
	if err != nil {
		err := fmt.Errorf("failed to create cipher: %v", err)
		global.Logger.Error(err)
		return "", err
	}

	// Initialize GCM mode for AES encryption
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		err = fmt.Errorf("failed to create GCM: %v", err)
		global.Logger.Error(err)
		return "", err
	}

	// Generate a random nonce for this encryption operation
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		err = fmt.Errorf("failed to generate nonce: %v", err)
		global.Logger.Error(err)
		return "", err
	}

	// Encrypt the plaintext using AES-GCM, sealing the nonce and plaintext together
	encryptedSecret := aesGCM.Seal(nonce, nonce, []byte(plainText), nil)

	// Return the result as a hex-encoded string
	return hex.EncodeToString(encryptedSecret), nil
}

// Decrypt decrypts an AES-256 GCM encrypted secret back to its original plain-text form.
//
// Parameters:
// - encryptedSecretHex: The hex-encoded string of the encrypted secret (produced by the Encrypt function).
// - masterKey: The same passphrase used for encryption, required to decrypt the secret.
//
// Returns:
// - The original plain-text secret (decrypted value).
// - An error if decryption fails, either due to an incorrect master key, tampering, or any decryption issue.
//
// Decryption Process:
// 1. The hex-encoded encrypted secret is decoded into a byte array.
// 2. The masterKey is hashed using SHA-256 to generate a 32-byte key for AES-256 decryption.
// 3. AES-256 GCM is initialized as the decryption mode.
// 4. The nonce is extracted from the encrypted data.
// 5. The remaining data (cipherText) is decrypted using the nonce and the hashed master key.
// 6. The decrypted plain-text is returned.
func Decrypt(encryptedSecretHex, masterKey string) (string, error) {
	// Decode the hex-encoded encrypted secret into a byte array
	encryptedSecret, err := hex.DecodeString(encryptedSecretHex)
	if err != nil {
		err = fmt.Errorf("failed to decode encrypted secret: %v", err)
		global.Logger.Error(err)
		return "", err
	}

	// Create a new AES cipher block using the hashed master key
	block, err := aes.NewCipher(createHash(masterKey))
	if err != nil {
		err = fmt.Errorf("failed to create cipher: %v", err)
		global.Logger.Error(err)
		return "", err
	}

	// Initialize GCM mode for AES decryption
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		err = fmt.Errorf("failed to create GCM: %v", err)
		global.Logger.Error(err)
		return "", err
	}

	// Extract the nonce from the encrypted data (the first part is the nonce)
	nonceSize := aesGCM.NonceSize()
	nonce, cipherText := encryptedSecret[:nonceSize], encryptedSecret[nonceSize:]

	// Decrypt the cipherText using the nonce and AES-GCM
	decryptedSecret, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		err = fmt.Errorf("failed to decrypt secret: %v", err)
		global.Logger.Error(err)
		return "", err
	}

	// Return the decrypted plain-text secret
	return string(decryptedSecret), nil
}

// createHash generates a SHA-256 hash of the provided master key (passphrase).
// AES-256 requires a 32-byte key, so the master key is hashed to ensure it meets this length requirement.
//
// Parameters:
// - key: The input string (masterKey) that is used to derive the encryption key.
//
// Returns:
// - A 32-byte slice representing the SHA-256 hash of the input key.
//
// Explanation:
// - AES-256 requires a key of exactly 32 bytes (256 bits).
// - The createHash function uses SHA-256 to produce a 32-byte key from the master key (passphrase), ensuring it meets the AES-256 key size requirement.
func createHash(key string) []byte {
	// Generate a SHA-256 hash of the key
	hash := sha256.Sum256([]byte(key))

	// Return the first 32 bytes of the hash (this is already 32 bytes because SHA-256 always produces a 256-bit output)
	return hash[:]
}
