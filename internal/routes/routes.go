package routes

import (
	"github.com/gin-gonic/gin"
)

//definindo rotas

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1") // base url
	{
		books := main.Group("books")
		{
			books.GET("/:id", controllers.ShowBook)
			books.GET("/", controllers.ShowBooks)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}
	return router
}
