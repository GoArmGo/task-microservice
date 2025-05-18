reset-db:
	docker-compose down -v
	docker-compose up -d

wait-db:
	@echo "⏳ Waiting for Postgres to be ready..."
	@./wait-for-it.sh localhost:5432 -t 30

migrate-up: wait-db
	migrate -path migrations -database "postgres://jt:secret@localhost:5432/tasks_db?sslmode=disable" up