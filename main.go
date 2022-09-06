package main

import (
	"exemplo/GO-API/handlers"
	"exemplo/GO-API/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Inicializando o servidor com Echo

	srvr := echo.New()
	
	srvr.Use(middleware.Logger())
	srvr.Use(middleware.Recover())

	srvr.Use(middleware.CORS())

	srvr.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"Access-Control-Allow-Origin", "*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	models.Connect()

	srvr.GET("/ListBooks", handlers.ListBooks)
	srvr.GET("/FindBook/:id", handlers.FindBook)
	srvr.POST("/CreateBooks/", handlers.CreateBook)
	srvr.PATCH("/UpdateBooks/:id", handlers.UpdateBook)
	srvr.DELETE("/books/:id", handlers.DeleteBook)

	srvr.Logger.Fatal(srvr.Start(":8080"))
}
