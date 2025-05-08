.DEFAULT_GOAL := help

ENV_FILE ?= .env
COMPOSE_FILE ?= infra/docker-compose/docker-compose.yaml
UI_API_DIR = services/ui-api

-include services/ui-api/Makefile

up-local: prepare-env-local ## Запуск контейнеров КРОМЕ UI-API
	docker-compose --env-file $(ENV_FILE) -f $(COMPOSE_FILE) up --build --scale ui-api=0 --scale searcher=0 --scale vectorizer=0

up: prepare-env ## Запуск всех контейнеров через docker-compose
	docker-compose --env-file $(ENV_FILE) -f $(COMPOSE_FILE) up --build

down: prepare-env ## Остановка всех контейнеров
	docker-compose --env-file $(ENV_FILE) -f $(COMPOSE_FILE) down

prepare-env-local:## Подготовить переменные окружения для локального запуска сервисов
	cp samples/.env.local .env

prepare-env: ## Подготовить переменные окружения для запуска сервисов в контейнерах
	cp samples/.env.docker .env

logs-ui-api: ## Просмотр логов только ui-api
	docker-compose --env-file $(ENV_FILE) -f $(COMPOSE_FILE) logs -f ui-api

help: ## Показывает список команд
	@awk -F':.*?## ' '/^[a-zA-Z0-9_-]+:.*## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
