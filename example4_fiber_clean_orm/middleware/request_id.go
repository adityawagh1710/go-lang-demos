package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func RequestID() fiber.Handler {
	return func(c fiber.Ctx) error {
		requestID := c.Get("X-Request-ID")

		if requestID == "" {
			requestID = uuid.New().String()
		}

		// Set in response
		c.Set("X-Request-ID", requestID)

		// Store in locals (accessible later)
		c.Locals("request_id", requestID)

		return c.Next()
	}
}
