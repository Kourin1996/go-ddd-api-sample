package books

import (
	"net/http"

	"github.com/Kourin1996/go-crud-api-sample/api/controllers/middleware"
	"github.com/Kourin1996/go-crud-api-sample/api/models/book"
	jwtToken "github.com/Kourin1996/go-crud-api-sample/api/models/jwt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

const BASE_PATH = "/books"

type BookHandler struct {
	bookService book.IBookService
}

func NewBookHandler(g *echo.Group, bookService book.IBookService) *BookHandler {
	handler := &BookHandler{bookService: bookService}

	publicGroup := g.Group(BASE_PATH)
	publicGroup.GET("/:hash_id", handler.GetBook)

	privateGroup := g.Group(BASE_PATH)
	privateGroup.Use(middleware.NewJwt())
	privateGroup.POST("", handler.PostBook)
	privateGroup.PUT("/:hash_id", handler.PutBook)
	privateGroup.DELETE("/:hash_id", handler.DeleteBook)

	return handler
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

func (h *BookHandler) PostBook(c echo.Context) error {
	tokenData := jwtToken.DecodeJWTToken(c.Get("user").(*jwt.Token))

	dto := &book.CreateBookDto{}
	if err := c.Bind(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	book, err := h.bookService.Create(tokenData, dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, book)
}

func (h *BookHandler) PutBook(c echo.Context) error {
	tokenData := jwtToken.DecodeJWTToken(c.Get("user").(*jwt.Token))

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

	book, err := h.bookService.Update(tokenData, hashId, dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, book)
}

func (h *BookHandler) DeleteBook(c echo.Context) error {
	tokenData := jwtToken.DecodeJWTToken(c.Get("user").(*jwt.Token))

	var hashId string
	err := echo.PathParamsBinder(c).String("hash_id", &hashId).BindError()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is not valid")
	}

	err = h.bookService.Delete(tokenData, hashId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.String(http.StatusOK, "")
}
