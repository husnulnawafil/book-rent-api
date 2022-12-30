package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/husnulnawafil/dot-id-task/models"
	"github.com/husnulnawafil/dot-id-task/modules"
	service "github.com/husnulnawafil/dot-id-task/services/book"
	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	bookService service.BookServiceInterface
}

func NewUserHandler(b service.BookServiceInterface) BookHandler {
	return BookHandler{
		bookService: b,
	}
}

func (b *BookHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var r *modules.Response
		var req models.Book
		err := c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, r.SendResponse("error_binding_data", http.StatusBadRequest, nil, nil))
		}

		if isValid := req.Name != "" &&
			req.Owner != 0 &&
			req.Author != "" &&
			req.Publisher != "" &&
			req.ISBN != ""; !isValid {
			return c.JSON(http.StatusBadRequest, r.SendResponse("please_fill_all_the_fields", http.StatusBadRequest, nil, nil))
		}

		book, code, err := b.bookService.Create(&req)
		if err != nil {
			return c.JSON(code, r.SendResponse(err.Error(), code, nil, nil))
		}

		return c.JSON(http.StatusCreated, r.SendResponse("success", http.StatusCreated, book, nil))
	}
}

func (b *BookHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		var r *modules.Response
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, r.SendResponse("id_must_be_an_integer", http.StatusBadRequest, nil, nil))
		}

		book, code, err := b.bookService.Get(uint(id))
		if err != nil {
			return c.JSON(code, r.SendResponse(err.Error(), code, nil, nil))
		}
		return c.JSON(http.StatusOK, r.SendResponse("success", http.StatusOK, book, nil))
	}
}

func (b *BookHandler) List() echo.HandlerFunc {
	return func(c echo.Context) error {
		var r *modules.Response
		limitQuery := c.QueryParam("limit")
		pageQuery := c.QueryParam("page")
		limit, err := strconv.Atoi(limitQuery)
		if err != nil && limitQuery != "" {
			return c.JSON(http.StatusBadRequest, r.SendResponse("limit_must_be_integer", http.StatusBadRequest, nil, nil))
		}
		page, err := strconv.Atoi(pageQuery)
		if err != nil && limitQuery != "" {
			return c.JSON(http.StatusBadRequest, r.SendResponse("page_must_be_integer", http.StatusBadRequest, nil, nil))
		}
		sort := c.QueryParam("sort")
		sortBy := c.QueryParam("sortBy")

		books, pgn, code, err := b.bookService.List(&modules.Pagination{
			Limit:  limit,
			Page:   page,
			Sort:   sort,
			SortBy: sortBy,
		})
		if err != nil {
			return c.JSON(code, r.SendResponse(err.Error(), code, nil, nil))
		}
		return c.JSON(http.StatusOK, r.SendResponse("success", http.StatusOK, books, pgn))
	}
}

func (b *BookHandler) Update() echo.HandlerFunc {
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

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, r.SendResponse("id_must_be_integer", http.StatusBadRequest, nil, nil))
		}

		book, code, err := b.bookService.Update(uint(id), req)
		if err != nil {
			return c.JSON(code, r.SendResponse(err.Error(), code, book, nil))
		}

		return c.JSON(http.StatusOK, r.SendResponse("success", http.StatusOK, book, nil))
	}
}

func (b *BookHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		var r *modules.Response

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, r.SendResponse("id_must_be_integer", http.StatusBadRequest, nil, nil))
		}

		if err = b.bookService.Delete(uint(id)); err != nil {
			return c.JSON(http.StatusInternalServerError, r.SendResponse(err.Error(), http.StatusInternalServerError, nil, nil))
		}

		return c.JSON(http.StatusOK, r.SendResponse("success", http.StatusOK, nil, nil))
	}
}
