package controllers

import (
	"strconv"
	"task-management-user-service/core/dto"
	"task-management-user-service/core/interfaces"
	"task-management-user-service/utils"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	services interfaces.UserService
}

func NewUserController(services interfaces.UserService) *UserController {
	return &UserController{services: services}
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {

	var req dto.ReqUser
	if err := ctx.BodyParser(&req); err != nil {
		return utils.HandleResponse(ctx, nil, "Invalid request data", err)
	}
	if err := c.services.CreateUser(req); err != nil {
		return utils.HandleResponse(ctx, nil, "Failed to create user", err)
	}
	return utils.HandleResponse(ctx, nil, "User created successfully", nil)
}

func (c *UserController) GetUsers(ctx *fiber.Ctx) error {
	users, err := c.services.GetUsers()
	if err != nil {
		return utils.HandleResponse(ctx, nil, "Failed to retrieve users", err)
	}
	return utils.HandleResponse(ctx, users, "Users retrieved successfully", nil)
}

func (c *UserController) GetUserByID(ctx *fiber.Ctx) error {

	userID := ctx.Params("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return utils.HandleResponse(ctx, nil, "Invalid user ID", err)
	}
	user, err := c.services.GetUserByID(uint(id))
	if err != nil {
		return utils.HandleResponse(ctx, nil, "Failed to retrieve user", err)
	}
	return utils.HandleResponse(ctx, user, "User retrieved successfully", nil)
}

func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
	var req dto.ReqUser
	if err := ctx.BodyParser(&req); err != nil {
		return utils.HandleResponse(ctx, nil, "Invalid request data", err)
	}
	userID := ctx.Params("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return utils.HandleResponse(ctx, nil, "Invalid user ID", err)
	}

	if err := c.services.UpdateUser(&req, id); err != nil {
		return utils.HandleResponse(ctx, nil, "Failed to update user", err)
	}
	return utils.HandleResponse(ctx, nil, "User updated successfully", nil)
}

func (c *UserController) DeleteUser(ctx *fiber.Ctx) error {

	var req struct {
		UserID string `json:"user_id" validate:"required"`
	}
	if err := ctx.BodyParser(&req); err != nil {
		return utils.HandleResponse(ctx, nil, "Invalid request data", err)
	}
	id, err := strconv.Atoi(req.UserID)
	if err != nil {
		return utils.HandleResponse(ctx, nil, "Invalid user ID", err)
	}
	if err := c.services.DeleteUser(uint(id)); err != nil {
		return utils.HandleResponse(ctx, nil, "Failed to delete user", err)
	}
	return utils.HandleResponse(ctx, nil, "User deleted successfully", nil)
}
