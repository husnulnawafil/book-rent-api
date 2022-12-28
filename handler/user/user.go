package handler

import (
	"net/http"

	service "github.com/husnulnawafil/dot-id-task/service/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService service.UserServiceInterface
}

func NewUserHandler(u service.UserServiceInterface) UserHandler {
	return UserHandler{
		userService: u,
	}
}

func (u *UserHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusCreated, nil)
	}
}
func (u *UserHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}
func (u *UserHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}
func (u *UserHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}
