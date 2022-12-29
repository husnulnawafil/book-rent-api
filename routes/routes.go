package routes

import (
	bh "github.com/husnulnawafil/dot-id-task/handlers/book"
	uh "github.com/husnulnawafil/dot-id-task/handlers/user"
	"github.com/labstack/echo/v4"
)

func UserPath(e *echo.Echo, uh *uh.UserHandler) {
	e.POST("/users", uh.Create())
	e.PATCH("/users", uh.Update())
	e.GET("/users/:id", uh.Get())
	e.DELETE("/users/:id", uh.Delete())
}
func BookPath(e *echo.Echo, bh *bh.BookHandler) {
	e.POST("/books", bh.Create())
	e.GET("/books/:id", bh.Get())
	e.PUT("/books/:id", bh.Update())
	e.DELETE("/books/:id", bh.Delete())
	e.GET("/books/", bh.List())
}
