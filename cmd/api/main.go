package main

import (
	"fmt"
	"os"

	"github.com/gentil-eilison/bougette-backend/cmd/api/handlers"
	"github.com/gentil-eilison/bougette-backend/cmd/api/middlewares"
	"github.com/gentil-eilison/bougette-backend/common"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Application struct {
	logger  echo.Logger
	server  *echo.Echo
	handler handlers.Handler
}

func main() {
	err := godotenv.Load()
	echoServer := echo.New()
	if err != nil {
		echoServer.Logger.Fatal("Error loading .env file")
	}

	db, err := common.NewPsql()

	if err != nil {
		echoServer.Logger.Fatal(err.Error())
	}

	app := Application{
		logger:  echoServer.Logger,
		server:  echoServer,
		handler: handlers.Handler{DB: db},
	}
	app.Routes()
	echoServer.Use(middlewares.AnotherMiddleWare, middlewares.CustomMiddleware, middleware.Logger())
	port := os.Getenv("APP_PORT")
	appAddress := fmt.Sprintf("localhost:%s", port)
	echoServer.Logger.Fatal(echoServer.Start(appAddress))
}
