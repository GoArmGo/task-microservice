package repository

import (
	"context"
	"task-microservice/internal/model"
)

// управляющий интерфейс для таск модели
type TaskRepository interface {
	Create(ctx context.Context, task *model.Task) error
	GetByID(ctx context.Context, id int64) (*model.Task, error)
	GetAll(ctx context.Context) ([]*model.Task, error)
	Update(ctx context.Context, task *model.Task) error
	Delete(ctx context.Context, id int64) error
	DeleteAll(ctx context.Context) error
}
