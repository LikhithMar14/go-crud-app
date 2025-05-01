package controllers

import (
	"net/http"
	"github.com/LikhithMar14/go-crud-app/config"
	"github.com/LikhithMar14/go-crud-app/models"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	result := config.DB.Find(&books)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	result := config.DB.First(&book, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}


func CreateBook(c *gin.Context) {
	var book models.Book


	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}


	c.JSON(http.StatusCreated, book)
}

func UpdateBook(c *gin.Context){
	var book models.Book
	id := c.Param("id")

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&models.Book{}).Where("id = ?", id).Updates(book)
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context){
	id := c.Param("id")
	config.DB.Delete(&models.Book{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
	
}