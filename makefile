APP_NAME=finbuddy
GO=go
DOCKER_COMPOSE=docker compose
DOCKER_FILE=Dockerfile
DOCKER_IMAGE=$(APP_NAME)
DOCKER_DEV_STAGE=dev

.PHONY: help build run test clean docker-build docker-up docker-down dev-up dev-down dev-shell lint fmt

# 🆘 Show help
help:
	@echo "Common commands:"
	@echo "  make build           Build Go binary locally"
	@echo "  make run             Run app locally (not in Docker)"
	@echo "  make test            Run Go tests"
	@echo "  make lint            Run go vet (basic lint)"
	@echo "  make fmt             Run go fmt"
	@echo "  make docker-build    Build Docker image for production"
	@echo "  make docker-up       Start containers (prod)"
	@echo "  make docker-down     Stop containers"
	@echo "  make dev-up          Start dev containers with hot reload"
	@echo "  make dev-down        Stop dev containers"
	@echo "  make dev-shell       Enter dev container shell"
	@echo "  make clean           Remove built binary"

# 🛠 Local build
build:
	$(GO) build -o $(APP_NAME) ./cmd/api

# ▶️ Run locally
run:
	./$(APP_NAME)

# 🧪 Test
test:
	$(GO) test ./...

# 🔍 Lint & Format
lint:
	$(GO) vet ./...

fmt:
	$(GO) fmt ./...

# 🧼 Clean up
clean:
	rm -f $(APP_NAME)

# 🐳 Production build
docker-build:
	docker build -t $(DOCKER_IMAGE) -f $(DOCKER_FILE) .

# 🐳 Start production-like environment
docker-up:
	$(DOCKER_COMPOSE) -f docker-compose.yml up --build

docker-down:
	$(DOCKER_COMPOSE) -f docker-compose.yml down

# 🧑‍💻 Start dev environment (with override + air)
dev-up:
	$(DOCKER_COMPOSE) up --build

dev-down:
	$(DOCKER_COMPOSE) down

dev-shell:
	$(DOCKER_COMPOSE) exec api sh

update-swagger:
	swag init -g cmd/api/main.go --output docs

reset-db:
	docker exec -i finbuddy-db psql -U finbuddy_user -d finbuddy_db -c "DROP SCHEMA public CASCADE; CREATE SCHEMA public;"
