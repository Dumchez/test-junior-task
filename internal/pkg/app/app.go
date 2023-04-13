package app

import (
	"fmt"
	"log"

	"github.com/Dumchez/test-junior-task/internal/app/endpoint"
	"github.com/Dumchez/test-junior-task/internal/app/mw"
	"github.com/Dumchez/test-junior-task/internal/app/service"
	"github.com/labstack/echo/v4"
)

type App struct {
	e      *endpoint.Endpoint
	s      *service.Service
	server *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.server = echo.New()

	a.server.GET("/status", a.e.Status, mw.RoleCheck)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("server is running")

	err := a.server.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
