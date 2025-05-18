DB_CONTAINER_NAME = task_pg
DB_NAME = tasks_db
DB_USER = jt
DB_PASS = secret
DB_PORT = 5432
DB_HOST = db # Важно: имя сервиса в docker-compose

reset-db:
	docker-compose down -v
	docker-compose up -d

wait-db:
	@echo "⏳ Waiting for Postgres to be ready..."
	@scripts/wait-for-it.sh $(DB_HOST):$(DB_PORT) -t 30

migrate-up:
	docker exec -e PGPASSWORD=$(DB_PASS) $(DB_CONTAINER_NAME) \
	  sh -c 'migrate -path /migrations -database "postgres://$(DB_USER):$(DB_PASS)@localhost:5432/$(DB_NAME)?sslmode=disable" up'

migrate-down:
	docker exec -e PGPASSWORD=$(DB_PASS) $(DB_CONTAINER_NAME) \
	  sh -c 'migrate -path /migrations -database "postgres://$(DB_USER):$(DB_PASS)@localhost:5432/$(DB_NAME)?sslmode=disable" down'

migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "⛔️ Usage: make migrate-create name=add_users_table"; \
	else \
		migrate create -ext sql -dir migrations $(name); \
	fi

psql:
	docker exec -it $(DB_CONTAINER_NAME) psql -U $(DB_USER) -d $(DB_NAME)

logs:
	docker logs -f $(DB_CONTAINER_NAME)