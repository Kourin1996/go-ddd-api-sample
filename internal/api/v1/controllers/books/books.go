package books

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/Kourin1996/go-crud-api-sample/pkg/models"
)

var books map[string]models.Book

func init() {
	books = make(map[string]models.Book)
	books["hoge"] = models.Book{
		Name:        "Hello Potter",
		Description: "Interestring Book",
		Price:       1500,
	}
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

	uuidObj, _ := uuid.NewUUID()
	id := uuidObj.String()
	books[id] = models.Book(*dto)

	return c.String(http.StatusCreated, id)
}

func getBook(c echo.Context) error {
	id := ""
	err := echo.PathParamsBinder(c).String("id", &id).BindError()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is not valid")
	}

	book, ok := books[id]
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, book)
}

func putBook(c echo.Context) error {
	id := ""
	err := echo.PathParamsBinder(c).String("id", &id).BindError()
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

	book := models.Book(*dto)
	books[id] = book

	return c.JSON(http.StatusOK, book)
}

func deleteBook(c echo.Context) error {
	id := ""
	err := echo.PathParamsBinder(c).String("id", &id).BindError()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is not valid")
	}

	_, ok := books[id]
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	delete(books, id)

	return c.String(http.StatusOK, "")
}
