package handlers

import (
	"net/http"

	"github.com/husnulnawafil/dot-id-task/models"
	"github.com/husnulnawafil/dot-id-task/modules"
	service "github.com/husnulnawafil/dot-id-task/services/user"
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
		var r *modules.Response
		var req models.User
		err := c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, r.SendResponse("error_binding_data", http.StatusBadRequest, nil, nil))
		}

		if isValid := req.Email != "" && req.FirstName != "" && req.LastName != "" && req.Phone != ""; !isValid {
			return c.JSON(http.StatusBadRequest, r.SendResponse("please_fill_all_the_fields", http.StatusBadRequest, nil, nil))
		}

		if ok := modules.ValidateEmail(req.Email); !ok {
			return c.JSON(http.StatusBadRequest, r.SendResponse("error_to_validate_email", http.StatusBadRequest, nil, nil))
		}

		phone, err := modules.ValidatePhone(req.Phone, "ID")
		if err != nil {
			return c.JSON(http.StatusBadRequest, r.SendResponse("error_to_validate_phone", http.StatusBadRequest, nil, nil))
		}
		req.Phone = phone

		user, code, err := u.userService.Create(&req)
		if err != nil {
			return c.JSON(code, r.SendResponse(err.Error(), code, nil, nil))
		}

		return c.JSON(http.StatusCreated, r.SendResponse("success", http.StatusCreated, user, nil))
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
