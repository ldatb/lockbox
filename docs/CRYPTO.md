### Cryptographic Operations

This project package implements encryption and decryption for sensitive data using AES-256 GCM (Galois/Counter Mode). AES-256 is a widely recognized, secure symmetric encryption algorithm, ensuring both confidentiality (encryption) and integrity (protection against data tampering).

#### 1. **Encryption Process (AES-256 GCM)**

The `Encrypt` function secures sensitive data by using a combination of AES-256 encryption and GCM mode to ensure both confidentiality and integrity. Here’s how it works:

##### Steps:
1. **Key Generation via SHA-256**:
   - The master passphrase (`masterKey`) is hashed using the SHA-256 algorithm to create a 32-byte key (required for AES-256).
   - AES-256 requires a 32-byte (256-bit) key, and hashing ensures that even if the passphrase is shorter or longer than 32 bytes, a secure, fixed-length key is generated.

2. **AES-256 GCM Mode**:
   - AES operates in **GCM mode** (Galois/Counter Mode), which is a secure encryption mode that provides both encryption (confidentiality) and message integrity (tamper detection).
   - AES-GCM requires an **initialization vector** or **nonce** for each encryption operation to ensure that the same data produces different encrypted outputs on subsequent encryptions (even if using the same key).

3. **Nonce Generation**:
   - A unique, random **nonce** (a number used once) is generated for every encryption operation. This prevents attackers from identifying patterns when the same plaintext is encrypted multiple times.

4. **Encryption**:
   - The plaintext (sensitive data) is encrypted with AES-256 GCM using the generated nonce and the 32-byte key derived from the master passphrase. The result is a **cipherText**, which contains both the encrypted data and the nonce.
   - The nonce is combined with the encrypted data and returned as a hex-encoded string.

##### Example Usage:
```go
encryptedData, err := secrets.Encrypt("my-secret-data", "my-master-key")
```

##### Security Considerations:
- **AES-256** is considered highly secure for most applications, providing 256 bits of encryption strength.
- The **GCM mode** not only encrypts but also ensures that any changes to the ciphertext (e.g., due to tampering) can be detected.
- **Nonce generation** is crucial to ensure that even identical data produces different ciphertexts.

#### 2. **Decryption Process (AES-256 GCM)**

The `Decrypt` function reverses the encryption, restoring the original plaintext from the encrypted data.

##### Steps:
1. **Hex Decoding**:
   - The hex-encoded encrypted data (containing both the nonce and the encrypted secret) is decoded back into a byte array.

2. **Key Generation via SHA-256**:
   - The same master passphrase used during encryption is hashed again using SHA-256 to regenerate the 32-byte key required for decryption.

3. **Nonce Extraction**:
   - The nonce, which is prepended to the encrypted secret, is extracted from the byte array. The remaining part is the actual encrypted data (cipherText).

4. **Decryption**:
   - AES-256 GCM uses the extracted nonce and the regenerated key to decrypt the ciphertext. If successful, the original plaintext (secret) is returned.

##### Example Usage:
```go
decryptedData, err := secrets.Decrypt(encryptedHexData, "my-master-key")
```

##### Security Considerations:
- Decryption will fail if the **nonce** or **ciphertext** is altered or if the wrong **master passphrase** is provided.
- AES-GCM ensures that even if data is tampered with, the decryption process will fail and prevent returning corrupted data.

#### 3. **Key Generation using SHA-256**

The `createHash` function plays a critical role in generating a consistent 32-byte key (required for AES-256) from any arbitrary-length passphrase.

##### How it Works:
- **SHA-256 Hashing**:
  - The master passphrase is hashed using **SHA-256**, which always produces a 32-byte (256-bit) output.
  - This ensures that the key passed to AES-256 is always of the correct length, regardless of the input passphrase length.

##### Example Usage:
```go
key := createHash("my-master-key")
```

##### Why Hash the Master Key?
- AES-256 requires a fixed-length key of 32 bytes. The `createHash` function ensures that the key is always 32 bytes by using SHA-256, no matter how long or short the master passphrase is.

### Summary of Security Features

- **AES-256 GCM**: 
  - Provides strong encryption and message integrity.
  - GCM (Galois/Counter Mode) ensures that any tampering with the ciphertext is detectable.

- **Nonce Usage**: 
  - A random nonce is used for every encryption operation, making the encryption unique for each operation, even when encrypting the same data.

- **SHA-256 for Key Derivation**: 
  - Ensures that the master passphrase is converted into a secure, fixed-length encryption key suitable for AES-256.

- **Hex Encoding**: 
  - The final encrypted result (including the nonce and the ciphertext) is returned as a hex-encoded string, making it easy to store or transmit.

By combining these features, the `secrets` package ensures that sensitive data is securely encrypted and decrypted, protecting it from unauthorized access or tampering.
