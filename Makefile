.DEFAULT_GOAL := help

MIGRATE := migrate
EXT := sql
DIR := db/migrations

define HELP_MESSAGE
Usage:
	make gen_migration FILENAME=create_users_table

Options:
	FILENAME: The name of the migration file to create (without extension)

Example:
	make gen_migration FILENAME=create_users_table
endef
export HELP_MESSAGE

help:
	@echo "$$HELP_MESSAGE"

gen_migration:
	$(MIGRATE) create -ext $(EXT) -dir $(DIR) -seq $(FILENAME)
