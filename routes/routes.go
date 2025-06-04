package routes

import (
	"github.com/Geek-Domain/book-collection-api/controllers"
	"github.com/Geek-Domain/book-collection-api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	auth := router.Group("/", middleware.AuthMiddleware())
	{
		auth.GET("/profile", func(c *gin.Context) {
			email := c.GetString("email")
			c.JSON(200, gin.H{"message": "You are authenticated", "email": email})
		})

		auth.POST("/books", controllers.CreateBook) // Create a new book
		auth.GET("/books", controllers.GetBooks)
		auth.GET("/books/:id", controllers.GetBookByID)
		auth.PUT("/books/:id", controllers.UpdateBookByID)
		auth.DELETE("/books/:id", controllers.DeleteBookByID)

	}
}
