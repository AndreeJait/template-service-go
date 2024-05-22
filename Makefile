export

.PHONY: build-http
build-http:
	@go build -o ./build/ ./app/http


.PHONY: run-http
run-http: build-http
	./build/http -E=DEV


.PHONY: migrate-new
migrate-new: ## create a new database migration
	@read -p "migration name: " name; \
	CGO_ENABLED="0" go install github.com/rubenv/sql-migrate/...@latest; \
	`go env GOPATH`/bin/sql-migrate new $${name}

.PHONY: migrate-up
migrate-up:
	@go run ./app/migration migrate up -E=DEV

.PHONY: migrate-down
migrate-down:
	@go run ./app/migration migrate down -E=DEV

.PHONY: migrate-fresh
migrate-fresh:
	@go run ./app/migration migrate fresh -E=DEV