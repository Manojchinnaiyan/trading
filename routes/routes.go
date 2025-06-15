package routes

import (
	"net/http"
	"trading-platform-backend/config"
	"trading-platform-backend/handlers"
	"trading-platform-backend/middleware"
	"trading-platform-backend/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, authService *services.AuthService, dataService *services.DataService, cfg *config.Config) {
	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	dataHandler := handlers.NewDataHandler(dataService)

	// Health check endpoint (open)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "healthy",
			"timestamp": "2024-01-01T00:00:00Z",
			"service":   "trading-platform-backend",
		})
	})

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// Authentication routes (no auth required)
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", authHandler.Signup)
			auth.POST("/login", authHandler.Login)
			auth.POST("/refresh", authHandler.RefreshToken)
		}

		// Protected routes (require JWT auth)
		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware(authService))
		{
			// Data endpoints as specified in the document
			protected.GET("/holdings", dataHandler.GetHoldings)
			protected.GET("/orderbook", dataHandler.GetOrderbook)
			protected.GET("/positions", dataHandler.GetPositions)
		}
	}

	// 404 handler
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Route not found",
		})
	})
}
