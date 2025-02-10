include .envrc
MIGRATION_PATH := ./cmd/migrate/migrations

.PHONY: migration
migration-%:
	@migrate create -seq -ext sql -dir $(MIGRATION_PATH) $*

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATION_PATH) -database="$(DB_ADDR)" up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATION_PATH) -database="$(DB_ADDR)" down

.PHONY: seed
seed:
	DB_ADDR=$(DB_ADDR) go run cmd/migrate/seed/main.go

# Add this to handle arguments
%:
	@: