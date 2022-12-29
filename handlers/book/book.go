package handlers

import (
	"net/http"

	service "github.com/husnulnawafil/dot-id-task/services/book"
	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	userService service.BookServiceInterface
}

func NewUserHandler(b service.BookServiceInterface) BookHandler {
	return BookHandler{
		userService: b,
	}
}

func (b *BookHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusCreated, nil)
	}
}

func (b *BookHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}

func (b *BookHandler) List() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}

func (b *BookHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}

func (b *BookHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}
