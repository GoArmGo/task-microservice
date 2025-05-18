reset-db:
	docker-compose down -v
	docker-compose up -d

DB_HOST ?= localhost

wait-db:
	@echo "⏳ Waiting for Postgres to be ready..."
	@scripts/wait-for-it.sh $(DB_HOST):5432 -t 30

migrate-up:
	@echo "⏳ Waiting for Postgres and tasks_db to be really ready..."
	@scripts/wait-for-it.sh $(DB_HOST):5432 -- \
	  sh -c 'until pg_isready -U jt -d tasks_db -h $(DB_HOST); do sleep 1; done && echo "✅ DB is ready!"; \
	         migrate -path migrations -database "postgres://jt:secret@$(DB_HOST):5432/tasks_db?sslmode=disable" up'

migrate-down:
	@scripts/wait-for-it.sh $(DB_HOST):5432 -- \
	  sh -c 'until pg_isready -U jt -d tasks_db -h $(DB_HOST); do sleep 1; done && echo "✅ DB is ready!"; \
	         migrate -path migrations -database "postgres://jt:secret@$(DB_HOST):5432/tasks_db?sslmode=disable" down'