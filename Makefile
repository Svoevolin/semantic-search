.DEFAULT_GOAL := help

-include services/api/Makefile

up: ## Запуск docker-compose
	docker-compose --env-file infra/docker-compose/.env -f infra/docker-compose/docker-compose.yaml up --build

down: ## Остановка docker-compose
	docker-compose --env-file infra/docker-compose/.env -f infra/docker-compose/docker-compose.yaml down

env: ## Копирует .env.sample → .env (если не существует)
	cp samples/.env .env

help: ## Показывает список команд
	@awk -F':.*?## ' '/^[a-zA-Z0-9_-]+:.*## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
