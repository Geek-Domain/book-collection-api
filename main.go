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
	if os.Getenv("GIN_MODE") != "release" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: .env file not loaded (not needed on Render)")
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
	router.Run(":" + port)

}
