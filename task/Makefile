include .env

MIGRATIONS_PATH = ./config/migrations

.PHONY: migrate-create
migrate-create:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(POSTGRESQL_URL) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(POSTGRESQL_URL) down $(filter-out $@,$(MAKECMDGOALS))

# Catch-all target to prevent "No rule to make target" errors when passing arguments
%:
	@: