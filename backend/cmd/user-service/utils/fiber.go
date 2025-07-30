package utils

import (
	"task-management-user-service/core/dto"
	"task-management-user-service/pkg"

	"github.com/gofiber/fiber/v2"
)

func HandleResponse(ctx *fiber.Ctx, data interface{}, mesSuccess string, err error) error {
	if err != nil {
		statusCode := fiber.StatusInternalServerError
		message := "An internal server error occurred."

		if appErr, ok := err.(*pkg.AppError); ok {
			statusCode = appErr.Code
			message = appErr.Message
		}

		return ctx.Status(statusCode).JSON(dto.Response{
			Status: false,
			Code:   statusCode,
			Msg:    message,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.Response{
		Status: true,
		Code:   fiber.StatusOK,
		Msg:    mesSuccess,
		Data:   data,
	})
}
