package handlers

import (
	"fmt"
	"net/http"
	"strconv"

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
		var r *modules.Response
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, r.SendResponse("id_must_be_an_integer", http.StatusBadRequest, nil, nil))
		}

		user, code, err := u.userService.Get(uint(id))
		if err != nil {
			return c.JSON(code, r.SendResponse(err.Error(), code, nil, nil))
		}
		return c.JSON(http.StatusOK, r.SendResponse("success", http.StatusOK, user, nil))
	}
}

func (u *UserHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var r *modules.Response

		req := map[string]interface{}{}
		c.Bind(&req)
		for i, v := range req {
			if ok := i == "id"; ok {
				delete(req, "id")
			}

			if len(req) == 0 {
				return c.JSON(http.StatusOK, r.SendResponse("nothing_to_be_updated", http.StatusOK, nil, nil))
			}

			if v == "" {
				return c.JSON(http.StatusBadRequest, r.SendResponse(fmt.Sprintf("%s_can_not_empty", i), http.StatusBadRequest, nil, nil))
			}
		}

		if _, ok := req["email"]; ok {
			if ok := modules.ValidateEmail(req["email"].(string)); !ok {
				return c.JSON(http.StatusBadRequest, r.SendResponse("error_to_validate_email", http.StatusBadRequest, nil, nil))
			}
		}

		if _, ok := req["email"]; ok {
			if ok := modules.ValidateEmail(req["email"].(string)); !ok {
				return c.JSON(http.StatusBadRequest, r.SendResponse("error_to_validate_email", http.StatusBadRequest, nil, nil))
			}
		}

		if _, ok := req["phone"]; ok {
			phone, err := modules.ValidatePhone(req["phone"].(string), "ID")
			if err != nil {
				return c.JSON(http.StatusBadRequest, r.SendResponse("error_to_validate_phone", http.StatusBadRequest, nil, nil))
			}
			req["phone"] = phone
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, r.SendResponse("id_must_be_integer", http.StatusBadRequest, nil, nil))
		}

		user, code, err := u.userService.Update(uint(id), req)
		if err != nil {
			return c.JSON(code, r.SendResponse(err.Error(), code, user, nil))
		}

		return c.JSON(http.StatusOK, r.SendResponse("success", http.StatusOK, user, nil))
	}
}

func (u *UserHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}
