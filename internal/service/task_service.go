package service

import (
	"context"
	"fmt"
	"task-microservice/internal/model"
	"task-microservice/internal/repository"
)

// слой с созданием тасков
type TaskService interface {
	CreateTask(ctx context.Context, task *model.Task) (*model.Task, error)
	GetTaskByID(ctx context.Context, id int64) (*model.Task, error)
	GetAllTasks(ctx context.Context) ([]*model.Task, error)
	UpdateTask(ctx context.Context, task *model.Task) (*model.Task, error)
	DeleteTask(ctx context.Context, id int64) error
	DeleteAllTasks(ctx context.Context) error
}

type TaskServiceImpl struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskServiceImpl {
	return &TaskServiceImpl{repo: repo}
}

func (s *TaskServiceImpl) CreateTask(ctx context.Context, task *model.Task) (*model.Task, error) {
	if task.Name == "" {
		return nil, fmt.Errorf("task name cannot be empty")
	}

	err := s.repo.Create(ctx, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskServiceImpl) GetTaskByID(ctx context.Context, id int64) (*model.Task, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *TaskServiceImpl) GetAllTasks(ctx context.Context) ([]*model.Task, error) {
	return s.repo.GetAll(ctx)
}

func (s *TaskServiceImpl) UpdateTask(ctx context.Context, task *model.Task) (*model.Task, error) {
	err := s.repo.Update(ctx, task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskServiceImpl) DeleteTask(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *TaskServiceImpl) DeleteAllTasks(ctx context.Context) error {
	return s.repo.DeleteAll(ctx)
}
