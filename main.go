package main

import (
	"exemplo/GO-API/controllers"
	"exemplo/GO-API/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Inicializando o servidor com Echo

	srvr := echo.New()

	srvr.Use(middleware.Logger())
	srvr.Use(middleware.Recover())

	models.Connect()

	srvr.GET("/books", controllers.FindBooks)
	srvr.GET("/books/:id", controllers.FindBook)
	srvr.POST("/books", controllers.CreateBook)
	srvr.PATCH("/books/:id", controllers.UpdateBook)
	srvr.DELETE("/books/:id", controllers.DeleteBook)

	srvr.Logger.Fatal(srvr.Start(":8080"))
}
