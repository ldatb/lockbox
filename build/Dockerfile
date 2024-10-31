# Define base images with explicit image versions and digests for reproducibility and security
ARG BUILDER_IMAGE=golang:1.32.2-alpine@sha256:9dd2625a1ff2859b8d8b01d8f7822c0f528942fe56cfe7a1e7c38d3b8d72d679
ARG RUNNER_IMAGE=alpine:3.20.3@sha256:beefdbd8a1da6d2915566fde36db9db0b524eb737fc57cd1367effd16dc0d06d

###################################
# STEP 1: Build executable binary #
###################################
FROM ${BUILDER_IMAGE} AS builder

# Install Git, SSL CA Certificates, and tzdata for time zones
# then clean up apk cache to minimize image size
RUN apk update && apk add --no-cache git ca-certificates tzdata && \
    update-ca-certificates && rm -rf /var/cache/apk/*

# Create a non-root user and group to avoid running as root, which is a security best practice
ENV USER=lockbox
ENV UID=1001
ENV GID=1001

# Create the user and group with no shell and no home directory, ensuring the user cannot log in
RUN addgroup --gid "${GID}" "${USER}" && \
    adduser --disabled-password --gecos "" --home "/nonexistent" --shell "/sbin/nologin" \
    --no-create-home --uid "${UID}" --ingroup "${USER}" "${USER}"

# Set the working directory to the Go project path within the container
WORKDIR $GOPATH/src/lockbox

# Copy Go module files first to cache dependencies, improving build performance for subsequent builds
COPY go.mod go.sum ./

# Download Go dependencies and verify checksums to ensure the integrity of downloaded dependencies
RUN go mod download && go mod verify

# Copy the entire source code into the container
COPY . .

# Build the Go binary with optimizations:
# - `CGO_ENABLED=0` disables CGo, producing a statically linked binary
# - `GOOS=linux` targets Linux as the OS
# - `GOARCH=amd64` targets 64-bit architecture
# - `-ldflags="-w -s"` removes debug information, reducing binary size
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/lockbox ./cmd/main.go

##################################
# STEP 2: Build a smaller image  #
##################################
FROM ${RUNNER_IMAGE} AS runner

# Set working directory for the application
WORKDIR /lockbox

# Install timezone data (`tzdata`) for correct time zone handling
RUN apk update && apk add --no-cache tzdata && rm -rf /var/cache/apk/*

# Copy time zone data and CA certificates from the builder stage, as they are necessary for SSL and correct time handling
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the built lockbox binary from the builder stage to the runner stage
COPY --from=builder /go/bin/lockbox .

# Create a non-root user in the final image with no login shell for enhanced security
# Use explicit UIDs and GIDs to ensure consistent behavior across systems
RUN echo 'lockbox:x:1001:1001::/nonexistent:/sbin/nologin' >> /etc/passwd && \
    echo 'lockbox:x:1001:' >> /etc/group

# Create a directory to store configuration files and set secure permissions
RUN mkdir /etc/lockbox && \
    chown -R lockbox:lockbox /lockbox && \
    chmod -R 755 /lockbox

# Run the container as the non-root `lockbox` user to minimize potential damage from vulnerabilities
USER lockbox:lockbox

# Define the entry point for the container to run the lockbox binary
ENTRYPOINT ["./lockbox"]
