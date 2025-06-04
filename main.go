package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Geek-Domain/book-collection-api/config"
	"github.com/Geek-Domain/book-collection-api/routes"
)

func main() {
	// Only try to load .env locally
	if os.Getenv("RENDER") != "true" {
		if err := godotenv.Load(); err != nil {
			log.Println("⚠️  .env file not found locally (ok if running on Render)")
		}
	}
	config.ConnectDB()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router)

	// Run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("✅ Server running on port", port)
	router.Run(":" + port)
}
