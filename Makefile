##############
# Parameters #
##############

# Go parameters
BINARY_NAME=lockbox
BUILD_DIR=./cmd

# Application parameters
CONFIG_FILE=lockbox.conf

# Docker parameters
DOCKER_IMAGE=lockbox
DOCKERFILE=build/Dockerfile
COMPOSE_FILE=deploy/docker-compose.yaml
COMPOSE_TEST_FILE=deploy/docker-compose-test.yaml

######
# Go #
######

# Download Go mods
.PHONY: mod
mod:
	go mod download
	go mod verify

.PHONY: dev
dev: mod
	go build -o $(BINARY_NAME) $(BUILD_DIR)/main.go

.PHONY: run
run:
	go run $(BUILD_DIR)/main.go --config-file $(CONFIG_FILE)

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	go clean
	rm -f $(BINARY_NAME)

##########
# Docker #
##########

.PHONY: dev-docker
docker-build:
	docker build -t $(DOCKER_IMAGE) -f $(DOCKERFILE) --network host .

.PHONY: docker-run
docker-run:
	docker run -d --network host --name ${BINARY_NAME} $(DOCKER_IMAGE)

.PHONY: docker-stop
docker-stop:
	docker stop $(DOCKER_IMAGE)
	docker rm $(DOCKER_IMAGE)

##################
# Docker Compose #
##################

.PHONY: compose-up
compose-up:
	docker compose -f $(COMPOSE_FILE) up -d

.PHONY: compose-down
compose-down:
	docker compose -f $(COMPOSE_FILE) down -v

###########
# Swagger #
###########

.PHONY: swagger-up
swagger-up:
	docker run -d -p 8080:8080 --network host -e SWAGGER_JSON=/lockbox/api/swagger.yaml -v .:/lockbox --name lockbox-swagger swaggerapi/swagger-ui

.PHONY: swagger-down
swagger-down:
	docker stop lockbox-swagger
	docker rm -f lockbox-swagger
