package books

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Kourin1996/go-crud-api-sample/api/models"
	"github.com/Kourin1996/go-crud-api-sample/api/services/book"
)

type BookHandler struct {
	bookService book.IBookService
}

func NewBookHandler(g *echo.Group, bookService book.IBookService) *BookHandler {
	handler := &BookHandler{bookService: bookService}

	group := g.Group("/books")
	group.POST("/", handler.postBook)
	group.GET("/:id", handler.getBook)
	group.PUT("/:id", handler.putBook)
	group.DELETE("/:id", handler.deleteBook)

	return handler
}

func (h *BookHandler) postBook(c echo.Context) error {
	dto := new(PostBookDto)
	if err := c.Bind(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	book := &models.Book{
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
	}
	book, err := h.bookService.CreateBook(book)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, *book)
}

func (h *BookHandler) getBook(c echo.Context) error {
	id := 0
	err := echo.PathParamsBinder(c).Int("id", &id).BindError()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is not valid")
	}

	book, err := h.bookService.GetBook(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, *book)
}

func (h *BookHandler) putBook(c echo.Context) error {
	id := 0
	err := echo.PathParamsBinder(c).Int("id", &id).BindError()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is not valid")
	}

	dto := new(PutBookDto)
	if err := c.Bind(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	book := &models.Book{
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
	}
	book, err = h.bookService.UpdateBook(id, book)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, *book)
}

func (h *BookHandler) deleteBook(c echo.Context) error {
	id := 0
	err := echo.PathParamsBinder(c).Int("id", &id).BindError()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is not valid")
	}

	err = h.bookService.DeleteBook(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.String(http.StatusOK, "")
}
