package middleware

import (
	"time"

	"example4_fiber_clean_orm/utils"

	"github.com/gofiber/fiber/v3"
)

func LoggerMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		requestID, _ := c.Locals("request_id").(string)

		utils.Log.Info().
			Str("request_id", requestID).
			Str("method", c.Method()).
			Str("path", c.Path()).
			Int("status", c.Response().StatusCode()).
			Dur("latency", time.Since(start)).
			Str("ip", c.IP()).
			Msg("request")

		return err
	}
}
