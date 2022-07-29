package main

import (
	"exemplo/GO-API/controllers"
	"exemplo/GO-API/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.Connect()

	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
