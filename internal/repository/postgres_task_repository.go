package repository

import (
	"context"
	"fmt"
	"task-microservice/internal/model"

	"github.com/jmoiron/sqlx"
)

// слой с бд
type PostgresTaskRepository struct {
	db *sqlx.DB
}

func NewPostgresTaskRepository(db *sqlx.DB) *PostgresTaskRepository {
	return &PostgresTaskRepository{db: db}
}

// создаем новый такс и кидаем в бд
func (r *PostgresTaskRepository) Create(ctx context.Context, task *model.Task) error {
	query := `
		INSERT INTO tasks (name, status, created_at, updated_at, due_date)
		VALUES (:name, :status, :created_at, :updated_at, :due_date)
		RETURNING id
	`
	// мапим поля запроса
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

// поиск по id
func (r *PostgresTaskRepository) GetByID(ctx context.Context, id int64) (*model.Task, error) {
	query := `
		SELECT id, name, status, created_at, updated_at, due_date
		FROM tasks
		WHERE id = $1
	`

	var task model.Task
	err := r.db.GetContext(ctx, &task, query, id)
	if err != nil {
		return nil, fmt.Errorf("get task by id: %w", err)
	}

	return &task, nil
}

// получение фулл таблицы
func (r *PostgresTaskRepository) GetAll(ctx context.Context) ([]*model.Task, error) {
	query := `
		SELECT id, name, status, created_at, updated_at, due_date
		FROM tasks
	`

	var tasks []*model.Task
	err := r.db.SelectContext(ctx, &tasks, query)
	if err != nil {
		return nil, fmt.Errorf("get all tasks: %w", err)
	}

	return tasks, nil
}

// обновить таску
func (r *PostgresTaskRepository) Update(ctx context.Context, task *model.Task) error {
	query := `
		UPDATE tasks
		SET name = :name, status = :status, updated_at = :updated_at, due_date = :due_date
		WHERE id = :id
	`

	_, err := r.db.NamedExecContext(ctx, query, task)
	if err != nil {
		return fmt.Errorf("update task: %w", err)
	}

	return nil
}

// удалить по id
func (r *PostgresTaskRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM tasks WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete task: %w", err)
	}

	return nil
}

// обнуление таблицы
func (r *PostgresTaskRepository) DeleteAll(ctx context.Context) error {
	query := `DELETE FROM tasks`

	_, err := r.db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to delete all tasks: %w", err)
	}

	return nil
}
