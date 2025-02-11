package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/koriebruh/simply_microservice/dto"
)

type Payload struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func WebResponse(ctx *fiber.Ctx, statusCode int, err error, message string, data interface{}) error {
	if err == nil {
		if data == nil {
			return ctx.Status(statusCode).JSON(Payload{
				Status:  "success",
				Message: message,
			})
		}
		return ctx.Status(statusCode).JSON(dto.WebResponse{
			Status:  "success",
			Message: message,
			Data:    data,
		})
	}

	return ctx.Status(statusCode).JSON(Payload{
		Status:  "error",
		Message: fmt.Sprintf("%s : %e", message, err),
	})
}
