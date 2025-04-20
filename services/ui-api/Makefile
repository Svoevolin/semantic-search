COMPOSE_FILE := infra/docker-compose/docker-compose.yaml

api-lint: ## Запустить golangci-lint
	golangci-lint run --config .golangci.yml ./...

swagger: ## Сгенерировать swagger для UI-API
	docker-compose -f $(COMPOSE_FILE) run --rm swagger-generator generate spec -i ./api/swagger.general.yaml -o ./api/swagger.json --scan-models

lint: ## Запустить golangci-lint
	golangci-lint run