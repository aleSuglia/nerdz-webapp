package handler

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"net/http"
)

type NerdzResponse struct {
	Data         interface{}
	Message      string
	HumanMessage string
}

var (
	e *echo.Echo = echo.New()
)

func welcome(c *echo.Context) {

	c.String(http.StatusOK, "Server is running!")
}

func Init() {
	//*************************//
	//   Built-in middleware   //
	//*************************//
	e.Use(mw.Logger)

	//************//
	//   Routes   //
	//************//
	//e.Post("/users", CreateUser)
	//e.Get("/users", GetUsers)

	e.Get("/users/:id/board", GetUserBoard)
	e.Get("/users/:id", GetUser)
	e.Get("/index", welcome)
	e.Get("/", welcome)
}

func Start() {
	e.Run(":8080")
}
