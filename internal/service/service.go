package service

import (
	"context"
	"fmt"
	"task-microservice/internal/model"
	"task-microservice/internal/repository"
)

type TaskService interface {
	Create(ctx context.Context, task *model.Task) error
	GetByID(ctx context.Context, id int64) (*model.Task, error)
	GetAll(ctx context.Context) ([]*model.Task, error)
	Update(ctx context.Context, task *model.Task) error
	Delete(ctx context.Context, id int64) error
}

type TaskServiceImpl struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskServiceImpl {
	return &TaskServiceImpl{repo: repo}
}

// CreateTask создаёт задачу через репозиторий
func (s *TaskServiceImpl) CreateTask(ctx context.Context, task *model.Task) error {
	if task.Name == "" {
		return fmt.Errorf("task name cannot be empty")
	}
	return s.repo.Create(ctx, task)
}

// GetTaskByID достаёт задачу по ID
func (s *TaskServiceImpl) GetTaskByID(ctx context.Context, id int64) (*model.Task, error) {
	return s.repo.GetByID(ctx, id)
}

// GetAllTasks возвращает все задачи
func (s *TaskServiceImpl) GetAllTasks(ctx context.Context) ([]*model.Task, error) {
	return s.repo.GetAll(ctx)
}
