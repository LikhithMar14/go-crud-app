package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/LikhithMar14/go-crud-app/controllers"
)

func SetupRoutes(router *gin.Engine){
	api := router.Group("/api")
	{
		books := api.Group("/books")
		{
			books.GET("", controllers.GetBooks)
			books.GET("/:id", controllers.GetBook)
			books.POST("", controllers.CreateBook)
			books.PUT("/:id", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}
}
