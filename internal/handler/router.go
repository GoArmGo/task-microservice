package handler

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(taskHandler *TaskHandler) *gin.Engine {
	r := gin.Default()

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
