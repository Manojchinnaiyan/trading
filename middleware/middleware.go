package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"trading-platform-backend/models"
	"trading-platform-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// Logger middleware
func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}

// Recovery middleware
func Recovery() gin.HandlerFunc {
	return gin.Recovery()
}

// JWT Auth middleware
func AuthMiddleware(authService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Error:   "Authorization header required",
				Message: "Please provide valid authorization token",
			})
			c.Abort()
			return
		}

		// Extract token from "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Error:   "Invalid authorization format",
				Message: "Use Bearer <token> format",
			})
			c.Abort()
			return
		}

		token := tokenParts[1]
		claims, err := authService.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Error:   "Invalid token",
				Message: err.Error(),
			})
			c.Abort()
			return
		}

		// Set user info in context
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Next()
	}
}

// Circuit Breaker middleware
func CircuitBreaker(cbService *services.CircuitBreakerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		breaker := cbService.GetBreaker("api")

		result, err := breaker.Execute(func() (interface{}, error) {
			c.Next()

			// Check if there was an error in the handler
			if len(c.Errors) > 0 {
				return nil, c.Errors[0]
			}

			// Check status code
			if c.Writer.Status() >= 500 {
				return nil, fmt.Errorf("server error: %d", c.Writer.Status())
			}

			return nil, nil
		})

		if err != nil {
			c.JSON(http.StatusServiceUnavailable, models.ErrorResponse{
				Error:   "Service temporarily unavailable",
				Message: "Please try again later",
			})
			c.Abort()
			return
		}

		_ = result // Unused in this case
	}
}

// Simple Rate Limit middleware for login route
func SimpleRateLimit(redisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		key := fmt.Sprintf("rate_limit:login:%s", clientIP)

		ctx := context.Background()

		// Get current count
		val, err := redisClient.Get(ctx, key).Result()
		if err != nil && err != redis.Nil {
			// If Redis fails, allow request to continue
			c.Next()
			return
		}

		count := 0
		if val != "" {
			count, _ = strconv.Atoi(val)
		}

		// Limit: 5 attempts per 15 minutes
		maxAttempts := 5
		if count >= maxAttempts {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error":   "Too many login attempts",
				"message": "Please try again in 15 minutes",
			})
			c.Abort()
			return
		}

		// Increment counter and set expiry
		pipe := redisClient.Pipeline()
		pipe.Incr(ctx, key)
		pipe.Expire(ctx, key, 15*time.Minute)
		pipe.Exec(ctx)

		c.Next()
	}
}
