package server

import (
	"task-management-task-service/core/controllers"

	"github.com/gofiber/fiber/v2"
)

func (s *server) registerRoutes(
	taskController *controllers.TaskController,
	router fiber.Router,
) {

	v1 := router.Group("/v1")

	v1.Post("/tasks", taskController.CreateTask)
	v1.Get("/tasks", taskController.GetTasks)
}
