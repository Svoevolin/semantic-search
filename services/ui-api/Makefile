COMPOSE_FILE := infra/docker-compose/docker-compose.yaml

swagger: ## Сгенерировать swagger для UI-API
	swag init -g cmd/ui-api/main.go -o api

lint: ## Запустить golangci-lint
	golangci-lint run