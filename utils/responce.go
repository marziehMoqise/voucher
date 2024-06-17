package utils

import (
	"github.com/gofiber/fiber/v2"
)

// Response Error
func ResponseError(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "error",
		"data": fiber.Map{
			"message": message,
		},
	})
}

// Response Error
func ResponseErrors(ctx *fiber.Ctx, statusCode int, message string) error {
	return ctx.Status(statusCode).JSON(fiber.Map{
		"status": "error",
		"data": fiber.Map{
			"message": message,
		},
	})
}
