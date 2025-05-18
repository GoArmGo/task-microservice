reset-db:
	docker-compose down -v
	docker-compose up -d

wait-db:
	@echo "⏳ Waiting for Postgres to be ready..."
	@scripts/wait-for-it.sh db:5432 -t 30

migrate-up:
	@echo "⏳ Waiting for Postgres and tasks_db to be really ready..."
	@scripts/wait-for-it.sh db:5432 -- \
	  sh -c 'until pg_isready -U jt -d tasks_db -h db; do sleep 1; done && echo "✅ DB is ready!"; \
	         migrate -path migrations -database "postgres://jt:secret@db:5432/tasks_db?sslmode=disable" up'