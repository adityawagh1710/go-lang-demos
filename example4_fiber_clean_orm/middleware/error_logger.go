package middleware

import (
	"example4_fiber_clean_orm/utils"

	"github.com/gofiber/fiber/v3"
)

func ErrorLogger() fiber.Handler {
	return func(c fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			requestID, _ := c.Locals("request_id").(string)

			utils.Log.Error().
				Str("request_id", requestID).
				Err(err).
				Msg("unhandled error")

			return err
		}
		return nil
	}
}
