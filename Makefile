ENV_FILE = ./config.env
include $(ENV_FILE)

docker-build:
	@docker compose build

docker-up:
	@docker compose up -d postgres
	@docker compose up -d url_shortener_app

docker-start:
	@docker compose start url_shortener_app

docker-stop:
	@docker compose stop url_shortener_app

docker-down:
	@docker compose down url_shortener_app
	@docker compose down postgres

goose-up:
	@goose $(GOOSE_DRIVER) $(GOOSE_DBSTRING) -dir $(GOOSE_MIGRATION_DIR) up

goose-down:
	@goose $(GOOSE_DRIVER) $(GOOSE_DBSTRING) -dir $(GOOSE_MIGRATION_DIR) down
