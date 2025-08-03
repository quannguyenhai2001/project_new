.PHONY: migrate-up migrate-down migrate-create migrate-version docker-up docker-down

DB_URL=postgres://admin:password123@localhost:5432/auth_db?sslmode=disable

# Database migrations
migrate-up:
	migrate -path ./migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path ./migrations -database "$(DB_URL)" down 1

migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations -seq $$name

migrate-version:
	migrate -path ./migrations -database "$(DB_URL)" version

# Docker commands
docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

# Server commands
run:
	go run cmd/server/main.go

build:
	go build -o bin/server cmd/server/main.go

start: build
	./bin/server

dev:
	air # hot reload (cần cài air: go install github.com/cosmtrek/air@latest)

# Combined commands
setup: docker-up migrate-up
	@echo "Database setup completed!"

reset: docker-down docker-up migrate-up
	@echo "Database reset completed!"

# Full development setup
dev-setup: setup run
	@echo "Development environment ready!"