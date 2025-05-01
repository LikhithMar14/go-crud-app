package routes

import (
	"github.com/LikhithMar14/go-crud-app/controllers"
	"github.com/LikhithMar14/go-crud-app/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Auth routes
	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	// Protected book routes
	books := router.Group("/books")
	books.Use(middleware.AuthMiddleware())
	{
		books.GET("", controllers.GetBooks)
		books.GET("/:id", controllers.GetBook)
		books.POST("", controllers.CreateBook)
		books.PATCH("/:id", controllers.UpdateBook)
		books.DELETE("/:id", controllers.DeleteBook)
	}
}
