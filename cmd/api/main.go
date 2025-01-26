package main

import (
	"liam-nak-na-api/internal/application"
	"liam-nak-na-api/internal/ports/http"

	"github.com/labstack/echo/v4"
)

func main() {
	triangleService := application.NewTriangleService()
	triangleHandler := http.NewTriangleHandler(triangleService)

	e := echo.New()

	e.POST("/triangle", triangleHandler.ClassifyTriangle)

	e.Logger.Fatal(e.Start(":8081"))
}
