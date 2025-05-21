package handler

import (
	"log"
	"task-microservice/internal/db"
	"task-microservice/internal/repository"
	"task-microservice/internal/service"
	"task-microservice/pkg/config"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// Загружаем конфиг
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Подключаем БД
	postgres := db.NewPostgres(cfg)
	if err := postgres.Ping(); err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}
	conn := postgres.DB

	// Инициализируем зависимости
	taskRepo := repository.NewPostgresTaskRepository(conn)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := NewTaskHandler(taskService)

	// Маршруты
	tasks := r.Group("/tasks")
	{
		tasks.POST("/", taskHandler.CreateTask)
		tasks.GET("/", taskHandler.GetAllTasks)
		tasks.GET("/:id", taskHandler.GetTaskByID)
		tasks.PUT("/:id", taskHandler.UpdateTask)
		tasks.DELETE("/:id", taskHandler.DeleteTask)
		tasks.DELETE("/", taskHandler.DeleteAllTasks)
	}

	return r
}
