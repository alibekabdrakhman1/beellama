APP_NAME = beellama
DB_URL = "postgres://postgres:password@localhost:5432/app_db?sslmode=disable"
MIGRATE = migrate -path migrations -database $(DB_URL)

.PHONY: all build  test migrations-up migrations-down

build:
	python3 generate_env.py
	docker compose up --build -d

test:
	go test ./... -v

migrations-up:
	$(MIGRATE) up

migrations-down:
	$(MIGRATE) up
