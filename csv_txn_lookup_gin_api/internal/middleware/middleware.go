package middleware

import (
	"csv-txn-lookup-gin-api/internal/logger"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.NewString()
		}
		c.Set("request_id", requestID)
		c.Header("X-Request-ID", requestID)
		c.Next()
	}
}

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			logger.Log.Error().
				Err(err).
				Str("path", c.Request.URL.Path).
				Msg("request failed")

			c.JSON(500, gin.H{"error": err.Error()})
		}
	}
}

func SimpleLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		reqID, _ := c.Get("request_id")

		logger.Log.Info().
			Int("status", status).
			Dur("latency", latency).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Interface("request_id", reqID).
			Msg("incoming request")
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Authorization header is required",
				"message": "Please add Bearer token",
			})
			return
		}

		const prefix = "Bearer "

		if len(authHeader) > len(prefix) && authHeader[:len(prefix)] == prefix {
			token := authHeader[len(prefix):]
			if token == "super-secret-token" {
				c.Set("user_id", "user-12345")
				c.Set("role", "admin")
				c.Set("is_authenticated", true)
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Invalid token",
			"message": "Please use 'super-secret-token'",
		})
	}
}
