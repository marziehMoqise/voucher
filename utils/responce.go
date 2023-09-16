package utils

import (
	"github.com/gofiber/fiber/v2"
)

// Response Error
func ResponseError(ctx *fiber.Ctx, message string) error{
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "error",
		"data": fiber.Map{
			"message": message,
		},
	})
}
