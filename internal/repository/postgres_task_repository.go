package repository

import (
	"context"
	"fmt"
	"task-microservice/internal/model"

	"github.com/jmoiron/sqlx"
)

type PostgresTaskRepository struct {
	db *sqlx.DB
}

func NewPostgresTaskRepository(db *sqlx.DB) *PostgresTaskRepository {
	return &PostgresTaskRepository{db: db}
}

// Create inserts a new task into the database
func (r *PostgresTaskRepository) Create(ctx context.Context, task *model.Task) error {
	query := `
		INSERT INTO tasks (name, status, created_at, updated_at, due_date)
		VALUES (:name, :status, :created_at, :updated_at, :due_date)
		RETURNING id
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return fmt.Errorf("prepare create: %w", err)
	}

	err = stmt.GetContext(ctx, &task.ID, task)
	if err != nil {
		return fmt.Errorf("insert task: %w", err)
	}

	return nil
}
