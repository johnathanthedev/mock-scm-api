.DEFAULT_GOAL := help

MIGRATE := migrate
EXT := sql
DIR := db/migrations

define HELP_MESSAGE
Usage:
	make gen_migration FILENAME=create_users_table
	make run_migrations
	make rollback

Options:
	FILENAME: The name of the migration file to create (without extension)

Example:
	make gen_migration FILENAME=create_users_table
	make run_migrations
	make rollback

endef
export HELP_MESSAGE

help:
	@echo "$$HELP_MESSAGE"

gen_migration:
	$(MIGRATE) create -ext $(EXT) -dir $(DIR) -seq $(FILENAME)
run_migrations:
	go run cmd/cli-tools/migrate/main.go
rollback:
	go run cmd/cli-tools/rollback/main.go