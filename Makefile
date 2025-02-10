# Makefile для создания миграций

# Переменные, которые будут использоваться в наших командах (Таргетах)
DB_DSN := "postgres://postgres:12345@localhost:5432/main?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Таргет для создания новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение миграций
migrate:
	$(MIGRATE) up

# Откат миграций
migrate-down:
	$(MIGRATE) down

# Команда для запуска приложения
run:
	go run cmd/app/main.go