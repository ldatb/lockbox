# How to Run XRS Lockbox Core Application

This guide explains how to build, run, and test the XRS Lockbox Core application using the provided `Makefile`. It includes instructions for building and running the application, running tests, building Docker images, and using Docker Compose to run services. Below are the detailed steps to get the application up and running.

## Prerequisites

Ensure you have the following installed on your system:
1. **Go** (version 1.16 or later)
2. **Docker**
3. **Docker Compose**

## Running the Application

### 1. Build the Application

Before running the application, you need to download the Go dependencies and build the binary.

Run the following command:

```bash
make build
```

This command will:
- Download the necessary Go modules.
- Compile the Go code into a binary named `lockbox` inside the project directory.

### 2. Run the Application

Once the binary is built, you can run the application with the provided configuration file (`lockbox.conf`).

Use this command to run the application:

```bash
make run
```

This will:
- Start the Lockbox Core application using the configuration file specified by `CONFIG_FILE` (`lockbox.conf`).
  
To specify a custom configuration file, modify the `CONFIG_FILE` variable in the `Makefile` or provide the path directly when running the application:

```bash
go run ./cmd/main.go --config-file <path-to-your-config-file>
```

### 3. Running Tests

You can run the unit tests using the following command:

```bash
make test
```

This will execute all the unit tests within the application and display the results.

### 4. Clean Up the Build

To clean up the generated binary and any temporary files:

```bash
make clean
```

This removes the binary file `lockbox` and cleans up the Go build environment.

## Using Docker

### 1. Build Docker Image

To build a Docker image for the Lockbox application, use the following command:

```bash
make docker-build
```

This command will:
- Build a Docker image named `lockbox` using the `Dockerfile` located in `build/Dockerfile`.

### 2. Run Docker Container

To run the Docker container after building the image:

```bash
make docker-run
```

This command will:
- Start the container using the `lockbox` Docker image, running the application within a container.

### 3. Stop and Remove Docker Container

To stop and remove the running container:

```bash
make docker-stop
```

This will stop and remove the `lockbox` container.

### 4. Rebuild and Recreate the Container

If you need to rebuild the image and recreate the container:

```bash
make docker-recreate
```

To stop, rebuild, and recreate the container:

```bash
make docker-recreate-stop
```

## Docker Compose

### 1. Start Services with Docker Compose

If you have a multi-container setup defined in `docker-compose.yaml`, you can start all the services using:

```bash
make compose-up
```

This will bring up all the services defined in the `deploy/docker-compose.yaml` file.

### 2. Stop Services with Docker Compose

To stop all running services:

```bash
make compose-down
```

This will bring down all the running containers and remove associated resources.

### 3. Recreate Docker Compose Services

To recreate the Docker Compose services, run:

```bash
make compose-recreate
```

This command will:
- Stop the services, rebuild the image, and bring the services back up.

### 4. Run Docker Compose for Testing

For testing purposes, you can start and stop services defined in `docker-compose-test.yaml`:

- To start testing services:

  ```bash
  make compose-test-up
  ```

- To stop testing services:

  ```bash
  make compose-test-down
  ```

### 5. Run Unit Tests After Compose

To run unit tests after bringing up the testing services:

```bash
make compose-test
```

This will:
- Bring down any running containers.
- Start the testing containers.
- Run the unit tests.
- Stop the testing containers after the tests.

## Swagger UI

### 1. Run Swagger UI

To run Swagger UI for API documentation, use the following command:

```bash
make swagger-up
```

This will:
- Start a Swagger UI container and expose it on port 8080, using the API definition in `/lockbox/api/swagger.yaml`.

### 2. Stop Swagger UI

To stop and remove the Swagger UI container:

```bash
make swagger-down
```

## Complete Build and Docker Workflow

To build both the Go binary and Docker image together:

```bash
make build-all
```

This command will:
- Build the Go binary.
- Build the Docker image for the application.

## Summary of Available `Makefile` Commands

|--------------------------|------------------------------------------------------------|
| Command                  | Description                                                |
|--------------------------|------------------------------------------------------------|
| `make build`             | Build the Go application.                                  |
| `make run`               | Run the application using the configuration file.          |
| `make test`              | Run all unit tests.                                        |
| `make clean`             | Clean up the build by removing the binary.                 |
| `make docker-build`      | Build the Docker image for the application.                |
| `make docker-run`        | Run the Docker container for the application.              |
| `make docker-stop`       | Stop and remove the running Docker container.              |
| `make docker-recreate`   | Rebuild the Docker image and recreate the container.       |
| `make compose-up`        | Start Docker Compose services.                             |
| `make compose-down`      | Stop Docker Compose services.                              |
| `make compose-test-up`   | Start Docker Compose services for testing.                 |
| `make compose-test-down` | Stop Docker Compose testing services.                      |
| `make swagger-up`        | Run Swagger UI for API documentation.                      |
| `make swagger-down`      | Stop and remove Swagger UI.                                |
| `make build-all`         | Build the Go binary and Docker image.                      |
|--------------------------|------------------------------------------------------------|

These commands offer a streamlined way to manage and run the Lockbox Core application across different environments, whether locally or in Docker.
