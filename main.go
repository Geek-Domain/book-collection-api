package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Geek-Domain/book-collection-api/config"
	"github.com/Geek-Domain/book-collection-api/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router)

	// Run server
	router.Run(":8080")
}
