package di

import (
	"task-microservice/internal/db"
	"task-microservice/internal/handler"
	"task-microservice/internal/repository"
	"task-microservice/internal/service"
	"task-microservice/pkg/config"

	"github.com/gin-gonic/gin"
)

type Container struct {
	Router *gin.Engine
}

func NewContainer() (*Container, error) {
	// 1. Загружаем конфигурацию
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	// 2. Инициализируем базу данных
	db := db.NewPostgres(cfg)
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// 3. Репозиторий
	taskRepo := repository.NewPostgresTaskRepository(db.DB)

	// 4. Сервис
	taskService := service.NewTaskService(taskRepo)

	// 5. Хендлер
	taskHandler := handler.NewTaskHandler(taskService)

	// 6. Роутер
	router := handler.NewRouter(taskHandler)

	return &Container{
		Router: router,
	}, nil
}
