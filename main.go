package main

import (
	"log"
	"os"

	"github.com/LikhithMar14/go-crud-app/config"
	"github.com/LikhithMar14/go-crud-app/models"
	"github.com/LikhithMar14/go-crud-app/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	if err:= godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	routes.SetupRoutes(router)

	config.ConnectDB()
	config.DB.AutoMigrate(&models.Book{})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}