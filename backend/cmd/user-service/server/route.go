package server

import (
	"task-management-user-service/core/controllers"

	"github.com/gofiber/fiber/v2"
)

func (s *server) registerRoutes(
	userController *controllers.UserController,
	router fiber.Router,
) {

	v1 := router.Group("/v1")

	v1.Post("/users", userController.CreateUser)
	v1.Get("/users", userController.GetUsers)
	v1.Get("/users/:id", userController.GetUserByID)
	v1.Put("/users/:id", userController.UpdateUser)
	v1.Delete("/users/:id", userController.DeleteUser)
}
