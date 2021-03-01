package books

import (
	"net/http"

	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"

	"github.com/Kourin1996/go-crud-api-sample/pkg/models"
	"github.com/Kourin1996/go-crud-api-sample/pkg/repositories/book"
)

var bookRepository book.IBookRepository

func init() {
	// todo: get from route
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "postgres",
		Database: "test",
	})
	bookRepository = book.NewRepository(db)
}

func SetRoutes(g *echo.Group) {
	booksGroup := g.Group("/books")

	booksGroup.POST("/", postBook)
	booksGroup.GET("/:id", getBook)
	booksGroup.PUT("/:id", putBook)
	booksGroup.DELETE("/:id", deleteBook)
}

func postBook(c echo.Context) error {
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
	err := bookRepository.CreateBook(book)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, *book)
}

func getBook(c echo.Context) error {
	id := 0
	err := echo.PathParamsBinder(c).Int("id", &id).BindError()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is not valid")
	}

	book, err := bookRepository.GetBook(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, *book)
}

func putBook(c echo.Context) error {
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
	err = bookRepository.UpdateBook(id, book)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, *book)
}

func deleteBook(c echo.Context) error {
	id := 0
	err := echo.PathParamsBinder(c).Int("id", &id).BindError()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is not valid")
	}

	err = bookRepository.DeleteBook(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.String(http.StatusOK, "")
}
