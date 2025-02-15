# Makefile для управления миграциями, генерацией кода и запуском приложения

# Переменные, которые будут использоваться в наших командах (таргетах)
DB_DSN := "postgres://postgres:12345@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Таргет для создания новой миграции
# Использование: make migrate-new NAME=<имя_миграции>
migrate-new:
ifndef NAME
	$(error NAME не указан. Используйте: make migrate-new NAME=<имя_миграции>)
endif
	migrate create -ext sql -dir ./migrations $(NAME)

# Применение миграций
migrate:
	$(MIGRATE) up

# Откат миграций
migrate-down:
	$(MIGRATE) down

# Запуск приложения
run:
	go run cmd/app/main.go

# Генерация кода из OpenAPI-спецификации
gen:
	mkdir -p ./internal/web/tasks
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go

# Установка зависимостей
deps:
	go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Проверка синтаксиса OpenAPI-спецификации
validate-openapi:
	docker run --rm -v $(PWD):/local openapitools/openapi-generator-cli validate -i /local/openapi/openapi.yaml

# Очистка сгенерированных файлов
clean:
	rm -f ./internal/web/tasks/api.gen.go

# Линтинг кода
lint:
	golangci-lint run --out-format=colored-line-number

# Помощь (список всех команд)
help:
	@echo "Доступные команды:"
	@echo "  make migrate-new NAME=<имя_миграции>  - Создать новую миграцию"
	@echo "  make migrate                          - Применить миграции"
	@echo "  make migrate-down                     - Откатить миграции"
	@echo "  make run                              - Запустить приложение"
	@echo "  make gen                              - Сгенерировать код из OpenAPI-спецификации"
	@echo "  make deps                             - Установить зависимости"
	@echo "  make validate-openapi                 - Проверить синтаксис OpenAPI-спецификации"
	@echo "  make clean                            - Очистить сгенерированные файлы"
	@echo "  make lint                             - Запустить статический анализ кода"
	@echo "  make help                             - Показать эту справку"