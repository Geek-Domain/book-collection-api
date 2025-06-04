package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Geek-Domain/book-collection-api/config"
	"github.com/Geek-Domain/book-collection-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var book models.Book
	bookCollection := config.GetCollection("books")

	// Get the authenticated user's email from JWT (set by middleware)
	email := c.GetString("email")
	book.UserEmail = email

	// Bind incoming JSON to the book struct
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Insert the book into the database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := bookCollection.InsertOne(ctx, book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Book created successfully",
		"book_id": result.InsertedID,
	})
}

func GetBooks(c *gin.Context) {
	bookCollection := config.GetCollection("books")
	email := c.GetString("email") // get user's email from token

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := bookCollection.Find(ctx, bson.M{"user_email": email})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}
	defer cursor.Close(ctx)

	var books []models.Book
	for cursor.Next(ctx) {
		var book models.Book
		if err := cursor.Decode(&book); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding book"})
			return
		}
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	bookCollection := config.GetCollection("books")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var book models.Book
	err = bookCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&book)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Make sure book belongs to logged-in user
	if book.UserEmail != c.GetString("email") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func UpdateBookByID(c *gin.Context) {
	id := c.Param("id")
	bookCollection := config.GetCollection("books")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	email := c.GetString("email")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Update only if the book belongs to the user
	result, err := bookCollection.UpdateOne(
		ctx,
		bson.M{"_id": objID, "user_email": email},
		bson.M{"$set": bson.M{
			"title":       book.Title,
			"author":      book.Author,
			"description": book.Description,
		}},
	)
	if err != nil || result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found or unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}

func DeleteBookByID(c *gin.Context) {
	id := c.Param("id")
	bookCollection := config.GetCollection("books")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	email := c.GetString("email")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := bookCollection.DeleteOne(ctx, bson.M{"_id": objID, "user_email": email})
	if err != nil || result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found or unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
