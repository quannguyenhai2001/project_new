package main

import (
	"backend/internal/handler"
	"backend/internal/middleware"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/pkg/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	cfg := config.Load()

	log.Printf("Starting server on port %s...", cfg.ServerPort)

	// Connect to database
	db := config.ConnectDB(cfg)
	defer db.Close()

	// Setup repositories
	userRepo := repository.NewUserRepository(db)

	// Setup services
	authService := service.NewAuthService(userRepo, cfg.JWTSecret)

	// Setup Gin router
	router := gin.Default()

	// Apply CORS middleware
	router.Use(middleware.CORSMiddleware())

	// Setup routes
	authHandler := handler.NewAuthHandler(authService)
	authHandler.SetupRoutes(router)

	// Add a simple health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Start the server
	log.Printf("Server running on http://localhost:%s", cfg.ServerPort)
	log.Printf("API endpoints:")
	log.Printf("  POST http://localhost:%s/api/auth/signup", cfg.ServerPort)
	log.Printf("  POST http://localhost:%s/api/auth/login", cfg.ServerPort)
	log.Printf("  GET http://localhost:%s/api/auth/me", cfg.ServerPort)
	log.Printf("  GET http://localhost:%s/health", cfg.ServerPort)

	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
