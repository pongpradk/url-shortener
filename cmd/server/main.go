package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pongpradk/url-shortener/internal/database"
	"github.com/pongpradk/url-shortener/internal/handler"
	"github.com/pongpradk/url-shortener/internal/repository"
	"github.com/pongpradk/url-shortener/internal/service"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Connect to database
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()
	log.Println("Database connected successfully!")

	// Initialize layers
	urlRepo := repository.NewURLRepository(db)
	urlService := service.NewURLService(urlRepo)
	urlHandler := handler.NewURLHandler(urlService)

	// Create Gin router
	router := gin.Default()

	// API routes
	api := router.Group("/api/v1")
	{
		api.POST("/data/shorten", urlHandler.HandleShorten)
	}

	// Redirect route
	router.GET("/:shortUrl", urlHandler.HandleRedirect)

	// Start server
	log.Println("Server starting on :8080")
	router.Run(":8080")
}
