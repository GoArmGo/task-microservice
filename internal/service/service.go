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

func (s *TaskServiceImpl) Create(ctx context.Context, task *model.Task) error {
	if task.Name == "" {
		return fmt.Errorf("task name cannot be empty")
	}
	return s.repo.Create(ctx, task)
}

func (s *TaskServiceImpl) GetByID(ctx context.Context, id int64) (*model.Task, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *TaskServiceImpl) GetAll(ctx context.Context) ([]*model.Task, error) {
	return s.repo.GetAll(ctx)
}

func (s *TaskServiceImpl) Update(ctx context.Context, task *model.Task) error {
	return s.repo.Update(ctx, task)
}

func (s *TaskServiceImpl) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
