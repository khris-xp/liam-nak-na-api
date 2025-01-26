package main

import (
	"liam-nak-na-api/internal/application"
	"liam-nak-na-api/internal/ports/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	triangleService := application.NewTriangleService()
	triangleHandler := http.NewTriangleHandler(triangleService)

	e := echo.New()

	// Add CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}))

	e.POST("/triangle", triangleHandler.ClassifyTriangle)

	e.Logger.Fatal(e.Start(":8081"))
}
