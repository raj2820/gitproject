package main

import (
	"log"
	"os"

	"gitlab-tool/internal/config"
	"gitlab-tool/internal/database"
	"gitlab-tool/internal/git"
	"gitlab-tool/internal/handlers"
	"gitlab-tool/internal/middleware"
	"gitlab-tool/internal/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := database.Init(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate database models
	if err := database.Migrate(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	repoRepo := repository.NewRepositoryRepository(db)

	// Initialize git service
	gitService := git.NewService(cfg.ReposPath)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(userRepo, cfg.JWTSecret)
	repoHandler := handlers.NewRepositoryHandler(repoRepo, gitService, cfg.ReposPath)
	healthHandler := handlers.NewHealthHandler()

	// Setup Gin router
	router := gin.Default()

	// Configure CORS middleware
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000", "http://127.0.0.1:3000"} // Nuxt dev server
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"}
	corsConfig.AllowCredentials = true
	corsConfig.ExposeHeaders = []string{"Content-Length"}

	// Apply CORS middleware
	router.Use(cors.New(corsConfig))

	// Serve static files
	router.Static("/static", "./static")
	router.LoadHTMLGlob("static/*.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// Health check endpoints
	router.GET("/health", healthHandler.HealthCheck)
	router.GET("/ready", healthHandler.ReadinessCheck)

	// Public routes
	router.POST("/auth/signup", authHandler.Signup)
	router.POST("/auth/login", authHandler.Login)

	// Protected routes
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	{
		// Repository routes
		protected.POST("/repos", repoHandler.CreateRepository)
		protected.GET("/repos", repoHandler.ListRepositories)
		protected.GET("/repos/:id", repoHandler.GetRepository)
		protected.DELETE("/repos/:id", repoHandler.DeleteRepository)

		// Git operations
		protected.POST("/repos/:id/clone", repoHandler.CloneRepository)
		protected.POST("/repos/:id/push", repoHandler.PushToRepository)
		protected.POST("/repos/:id/pull", repoHandler.PullFromRepository)
	}

	// Git HTTP backend routes (for git clone/push/pull)
	gitGroup := router.Group("/git")
	{
		// Use wildcard routing to capture all Git operations
		gitGroup.Any("/:username/:repo/*action", repoHandler.GitHTTPBackend)
	}

	// Create repos directory if it doesn't exist
	if err := os.MkdirAll(cfg.ReposPath, 0755); err != nil {
		log.Fatalf("Failed to create repos directory: %v", err)
	}

	log.Printf("Starting server on port %s", cfg.Port)
	log.Printf("Repositories will be stored in: %s", cfg.ReposPath)
	log.Printf("Health check available at: http://localhost:%s/health", cfg.Port)
	log.Printf("CORS enabled for Nuxt frontend (http://localhost:3000)")

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
