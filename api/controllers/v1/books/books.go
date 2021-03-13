package books

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Kourin1996/go-crud-api-sample/api/models/book"
)

type BookHandler struct {
	bookService book.IBookService
}

func NewBookHandler(g *echo.Group, bookService book.IBookService) *BookHandler {
	handler := &BookHandler{bookService: bookService}

	group := g.Group("/books")
	group.POST("", handler.PostBook)
	group.GET("/:hash_id", handler.GetBook)
	group.PUT("/:hash_id", handler.PutBook)
	group.DELETE("/:hash_id", handler.DeleteBook)

	return handler
}

func (h *BookHandler) PostBook(c echo.Context) error {
	dto := &book.CreateBookDto{}
	if err := c.Bind(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	book, err := h.bookService.Create(dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, book)
}

func (h *BookHandler) GetBook(c echo.Context) error {
	var hashId string
	err := echo.PathParamsBinder(c).String("hash_id", &hashId).BindError()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "hashId is not valid")
	}

	book, err := h.bookService.Get(hashId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, book)
}

func (h *BookHandler) PutBook(c echo.Context) error {
	var hashId string
	err := echo.PathParamsBinder(c).String("hash_id", &hashId).BindError()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is not valid")
	}

	dto := &book.UpdateBookDto{}
	if err := c.Bind(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	book, err := h.bookService.Update(hashId, dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, book)
}

func (h *BookHandler) DeleteBook(c echo.Context) error {
	var hashId string
	err := echo.PathParamsBinder(c).String("hash_id", &hashId).BindError()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is not valid")
	}

	err = h.bookService.Delete(hashId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.String(http.StatusOK, "")
}
