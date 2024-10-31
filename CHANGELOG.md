# Changelog

## [v1.0.0] - 2024-10-23

### Added

- **AES-256 GCM Encryption**:
  - Secure encryption of sensitive data using AES-256 GCM, ensuring both confidentiality and integrity.
  - Each encryption operation generates a unique nonce, preventing data pattern detection.

- **Decryption Using AES-256 GCM**:
  - Decrypts data encrypted with AES-256 GCM back to its original plaintext using the master passphrase.

- **SHA-256 Key Derivation**:
  - Passphrases are hashed using SHA-256 to generate the 32-byte keys required for AES-256 encryption.

- **Environment-Based Cryptographic Passphrase**:
  - `MASTER_CRYPTO_PASS` environment variable support for managing the cryptographic passphrase.
  - If not set, a random passphrase is generated, with warnings to avoid using random keys in production.

- **Hex Encoding for Encrypted Data**:
  - Encrypted data, along with the nonce, is returned as a hex-encoded string for easy storage and transmission.

- **RESTful API for Secure Data Management**:
  - Introduced an API for securely managing data within Lockbox.
  - Endpoints allow the creation, retrieval, and deletion of encrypted secrets.

- **HTTP Server Configuration**:
  - Configurable via the `lockbox.conf` file, allowing custom host and port settings for the API server.
  - Provides secure, centralized access to encrypted data through HTTP requests.

- **API Routing with HTTP Handlers**:
  - Integrated API router for handling incoming HTTP requests.
  - Support for custom middleware and request handling, allowing easy expansion of the API.

- **Customizable Logging**:
  - Configurable logging levels (debug, info, warn, error) and file paths via the `lockbox.conf` file.
  - Logs API requests and responses for easier troubleshooting and auditing.

- **Makefile Automation**:
  - Added commands for building, running, testing, and cleaning the application.
  - Docker build, run, and Docker Compose commands for deployment.
  - Swagger UI Docker commands to access API documentation.
