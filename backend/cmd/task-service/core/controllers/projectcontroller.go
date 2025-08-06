package controllers

import (
	"task-management-task-service/core/dto"
	"task-management-task-service/core/interfaces"
	"task-management-task-service/utils"

	"github.com/gofiber/fiber/v2"
)

type ProjectController struct {
	services interfaces.ProjectService
}

func NewProjectController(services interfaces.ProjectService) *ProjectController {
	return &ProjectController{services: services}
}

func (c *ProjectController) CreateProject(ctx *fiber.Ctx) error {
	var req dto.CreateProjectRequest
	if err := ctx.BodyParser(&req); err != nil {
		return utils.HandleResponse(ctx, nil, "Invalid request data", err)
	}
	if err := c.services.CreateProject(req); err != nil {
		return utils.HandleResponse(ctx, nil, "Failed to create project", err)
	}
	return utils.HandleResponse(ctx, nil, "Project created successfully", nil)
}

func (c *ProjectController) GetProjects(ctx *fiber.Ctx) error {
	projects, err := c.services.GetProjects()
	if err != nil {
		return utils.HandleResponse(ctx, nil, "Failed to retrieve projects", err)
	}
	return utils.HandleResponse(ctx, projects, "Projects retrieved successfully", nil)
}
