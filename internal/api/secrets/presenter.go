package secrets

// SecretResponseUUID represents the structure of a secret containing the UUID.
type SecretResponseUUID struct {
	// ID is the UUId associated with the secret
	ID string `json:"id"`

	// Key is the unique identifier or key associated with the secret.
	Key string `json:"key"`
}

// SecretResponsePlain represents the structure of a secret without encryption.
// This struct is used to present the decrypted secret back to the client in a safe and structured format.
type SecretResponsePlain struct {
	// Key is the unique identifier or key associated with the secret.
	Key string `json:"key"`

	// Value is the decrypted value of the secret.
	// This represents the actual sensitive information that was previously encrypted and is now being returned in plain text.
	Value string `json:"value"`
}
