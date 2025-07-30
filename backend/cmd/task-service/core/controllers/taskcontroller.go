package controllers

import (
	"fmt"
	"task-management-task-service/core/dto"
	"task-management-task-service/core/interfaces"
	"task-management-task-service/utils"

	"github.com/gofiber/fiber/v2"
)

type TaskController struct {
	services interfaces.TaskService
}

func NewTaskController(services interfaces.TaskService) *TaskController {
	return &TaskController{services: services}
}

func (c *TaskController) CreateTask(ctx *fiber.Ctx) error {

	var req dto.ReqTask
	if err := ctx.BodyParser(&req); err != nil {
		return utils.HandleResponse(ctx, nil, "Invalid request data", err)
	}
	fmt.Println("Received request to create task:", req)
	if err := c.services.CreateTask(req); err != nil {
		return utils.HandleResponse(ctx, nil, "Failed to create task", err)
	}
	fmt.Println("Task created successfully:", req)
	return utils.HandleResponse(ctx, nil, "Task created successfully", nil)
}

func (c *TaskController) GetTasks(ctx *fiber.Ctx) error {
	tasks, err := c.services.GetTasks()
	if err != nil {
		return utils.HandleResponse(ctx, nil, "Failed to retrieve tasks", err)
	}
	return utils.HandleResponse(ctx, tasks, "Tasks retrieved successfully", nil)
}
