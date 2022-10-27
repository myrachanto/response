package routes

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/myrachanto/respon/src/api/load"
)

func ApiLoader() {
	loader := load.NewloadController(load.NewloadService(load.NewloadRepo()))
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	api := e.Group("/api")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file in routes")
	}
	api.GET("/getURL", loader.GetURL).Name = "get-url"

	PORT := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(PORT))
}
