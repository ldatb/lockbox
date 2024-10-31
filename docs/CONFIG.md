# XRS Lockbox Core Configuration File Guide

## Overview

The XRS Lockbox Core configuration file (`.conf`) is used to define key settings related to the server and logging behavior of the application. This file must follow the standard INI-like format, where sections are enclosed in square brackets (`[ ]`), and key-value pairs are provided for specific configurations within each section. 

This document explains the required sections and configuration options for the `.conf` file used by Lockbox.

### Structure of the `.conf` File

The configuration file is divided into sections. Each section addresses a specific aspect of the application. Below are the sections and their respective keys.

#### [server] Section

The `[server]` section defines the basic network configurations for the application. It contains the following key-value pairs:

- **port**: Defines the port number on which the server will run.
  - Example: `port = 8080`
  - Type: Integer
  - Default: `8080`

- **host**: Specifies the host IP address or domain name for the server.
  - Example: `host = 0.0.0.0`
  - Type: String
  - Default: `0.0.0.0` (This means the server will be accessible from all network interfaces)

##### Example:

```conf
[server]
port = 8080
host = 0.0.0.0
```

#### [database] Section

The `[database]` section configures the database connection for Lockbox. It contains the following key-value pairs:

- **host**: The hostname or IP address of the database server.
  - Example: `host = localhost`
  - Type: String
  - Default: `localhost`

- **port**: The port number on which the database is listening.
  - Example: `port = 5432`
  - Type: Integer
  - Default: `5432`

- **username**: The username used to connect to the database.
  - Example: `username = lockboxuser`
  - Type: String
  - Default: `lockboxuser`

- **password**: The password for the specified database user.
  - Example: `password = dbpass`
  - Type: String
  - Default: `password`

- **db_name**: The name of the database to connect to.
  - Example: `db_name = lockboxdb`
  - Type: String
  - Default: `lockboxdb`

- **ssl_mode**: The SSL mode for database connections (e.g., disable, require, etc.).
  - Example: `ssl_mode = disable`
  - Type: String
  - Default: `disable`

- **max_idle_conns**: Maximum number of idle connections in the connection pool.
  - Example: `max_idle_conns = 10`
  - Type: Integer
  - Default: `10`

- **max_open_conns**: Maximum number of open connections to the database.
  - Example: `max_open_conns = 100`
  - Type: Integer
  - Default: `100`

- **max_conn_life**: Maximum lifetime (in seconds) for each connection in the pool.
  - Example: `max_conn_life = 60`
  - Type: Integer
  - Default: `60`

##### Example:

```conf
[database]
host = localhost
port = 5432
username = lockboxuser
password = dbpass
db_name = lockboxdb
ssl_mode = disable
max_idle_conns = 10
max_open_conns = 100
max_conn_life = 60
```

#### [logging] Section

The `[logging]` section configures how the application handles logging. This helps in troubleshooting, auditing, and monitoring the system's behavior. It includes the following key-value pairs:

- **level**: Specifies the log level. This defines the severity of logs to be captured. Common values include:
  - `debug`: Detailed debugging information.
  - `info`: General operational messages.
  - `warn`: Warning messages about potential issues.
  - `error`: Error messages indicating failures.
  
  - Example: `level = info`
  - Type: String
  - Default: `info`

- **filepath**: Defines the path to the log file where logs will be written.
  - Example: `filepath = Lockbox.log`
  - Type: String
  - Default: `Lockbox.log`

##### Example:

```conf
[logging]
level = info
filepath = Lockbox.log
```

### Example Configuration File

Here is a complete example of how your `config.conf` file should look (this is an example file for testing purposes):

```conf
[server]
host = 0.0.0.0
port = 8080

[database]
host = postgres
port = 5432
username = lockboxuser
password = password
db_name = lockboxdb
ssl_mode = disable
max_idle_conns = 10
max_open_conns = 100
max_conn_life = 60

[logging]
level = info
filepath = lockbox.log
```

### Cryptographic Passphrase Management

Lockbox uses a cryptographic passphrase for securing sensitive data. The passphrase is loaded from the environment and can be dynamically generated if not provided. This section explains how Lockbox handles the cryptographic passphrase using the environment variable MASTER_CRYPTO_PASS.

A environment variable named **MASTER_CRYPTO_PASS** should be set in production environments to ensure consistent encryption and decryption of sensitive data. If the environment variable is not set, the application will generate a random passphrase, which is **NOT suitable for production** as it will change each time the application is restarted, leading to data inconsistency.

### Requirements

1. **File Format**: 
   - The configuration file must be in the `.conf` format with sections enclosed in square brackets (`[ ]`), and keys mapped to values using `=`.

2. **Default Values**: 
   - If any key is missing, Lockbox may use its default values as described above. It is recommended to explicitly define these values to avoid unexpected behaviors.

3. **Placement**: 
   - Ensure the configuration file is accessible to the application at runtime. The path to the configuration file can be provided using the `--config-file` flag when running the application.

### Common Errors

- **Invalid Section Names**: Ensure that section names are enclosed in square brackets (`[ ]`).
- **Incorrect Key-Value Syntax**: Always use the `=` sign between keys and values, with no spaces between the key and the `=`.

### Additional Notes

- **Modifying the File**: You can edit the `.conf` file at any time to update the server or logging configurations. However, changes will take effect only after restarting the Lockbox application.
- **Cryptographic Key:** Ensure that the cryptographic passphrase `MASTER_CRYPTO_PASS` is properly set in production environments to avoid issues with encryption and decryption of data.
