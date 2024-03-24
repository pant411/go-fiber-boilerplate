package errorHandler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func AllException(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError
	var Message string = "error"

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		Message = e.Message
	}

	// Send custom error page
	if err != nil {
		// In case the SendFile fails
		return ctx.Status(code).JSON(GlobalErrorHandlerResp{
			Success: false,
			Message: Message,
		})
	}

	// Return from handler
	return nil
}
