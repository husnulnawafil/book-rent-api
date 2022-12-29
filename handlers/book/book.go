package handlers

import (
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
			req.ISBN != "" &&
			req.Price != 0; !isValid {
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
		limit, err := strconv.Atoi(c.QueryParam("limit"))
		if err != nil {
			limit = 0
		}
		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil {
			page = 0
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
		return c.JSON(http.StatusOK, nil)
	}
}

func (b *BookHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}
