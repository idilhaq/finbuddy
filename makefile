APP_NAME=finbuddy
DOCKER_COMPOSE=docker-compose
GO=go

.PHONY: help build run test clean docker-build docker-up docker-down

help:
	@echo "Common commands:"
	@echo "  make build          Build Go binary locally"
	@echo "  make run            Run app locally (not in Docker)"
	@echo "  make test           Run Go tests"
	@echo "  make docker-build   Build Docker image"
	@echo "  make docker-up      Start containers via docker-compose"
	@echo "  make docker-down    Stop containers"
	@echo "  make clean          Remove built binaries"

build:
	$(GO) build -o $(APP_NAME) .

run:
	./$(APP_NAME)

test:
	$(GO) test ./...

clean:
	rm -f $(APP_NAME)

docker-build:
	docker build -t $(APP_NAME) .

docker-up:
	$(DOCKER_COMPOSE) up --build

docker-down:
	$(DOCKER_COMPOSE) down
