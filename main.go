package main

import (
	"rest-echo/controllers"
	"rest-echo/models"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/users/:id", controllers.GetUser)
	e.POST("/users", controllers.Save)

	models.InitDB()

	// Run Server
	e.Logger.Fatal(e.Start(":8000"))
}
