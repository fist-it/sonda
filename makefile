# Makefile for Sonda Service Monitor

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=sonda
BINARY_UNIX=$(BINARY_NAME)_unix
MAIN_PATH=./cmd/sonda

# Build directory
BUILD_DIR=build

# Docker parameters
DOCKER_COMPOSE=docker-compose
DOCKERFILE_PATH=./deploy/Dockerfile
DOCKER_IMAGE_NAME=sonda
DOCKER_TAG=latest

.PHONY: all build clean test coverage lint deps tidy run docker-build docker-run docker-compose-up docker-compose-down help

all: test build

build:
	mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)

clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

test:
	$(GOTEST) -v ./...

coverage:
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

lint:
	golangci-lint run

deps:
	$(GOGET) -u ./...

tidy:
	$(GOMOD) tidy

run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_TAG) -f $(DOCKERFILE_PATH) .

docker-run: docker-build
	docker run --rm $(DOCKER_IMAGE_NAME):$(DOCKER_TAG)

docker-compose-up:
	cd deploy && $(DOCKER_COMPOSE) up -d

docker-compose-down:
	cd deploy && $(DOCKER_COMPOSE) down

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_UNIX) $(MAIN_PATH)

help:
	@echo "Available targets:"
	@echo "  all            - Run tests and build binary"
	@echo "  build          - Build the project binary"
	@echo "  clean          - Remove build artifacts"
	@echo "  test           - Run all tests"
	@echo "  coverage       - Generate code coverage report"
	@echo "  lint           - Run linter"
	@echo "  deps           - Update dependencies"
	@echo "  tidy           - Clean up go.mod and go.sum"
	@echo "  run            - Build and run the binary"
	@echo "  docker-build   - Build Docker image"
	@echo "  docker-run     - Run Docker container"
	@echo "  docker-compose-up - Start all services with Docker Compose"
	@echo "  docker-compose-down - Stop all services with Docker Compose"
	@echo "  build-linux    - Cross compile for Linux"

