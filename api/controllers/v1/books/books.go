package books

import (
	"net/http"

	"github.com/Kourin1996/go-crud-api-sample/api/common"
	"github.com/Kourin1996/go-crud-api-sample/api/controllers/middleware"
	"github.com/Kourin1996/go-crud-api-sample/api/models/book"
	"github.com/Kourin1996/go-crud-api-sample/api/models/errors"
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

	privateGroup := g.Group(BASE_PATH)
	privateGroup.Use(middleware.NewJwt())
	privateGroup.POST("", handler.PostBook)
	privateGroup.PUT("/:hash_id", handler.PutBook)
	privateGroup.DELETE("/:hash_id", handler.DeleteBook)

	publicGroup := g.Group(BASE_PATH)
	publicGroup.GET("", handler.GetBooks)
	publicGroup.GET("/:hash_id", handler.GetBook)

	return handler
}

func (h *BookHandler) GetBooks(c echo.Context) error {
	dto := &book.GetBooksDto{}
	if err := common.BindAndValidate(c, dto); err != nil {
		return errors.NewInvalidRequestError(err)
	}

	res, err := h.bookService.GetBooks(dto)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (h *BookHandler) GetBook(c echo.Context) error {
	var hashId string
	err := echo.PathParamsBinder(c).String("hash_id", &hashId).BindError()
	if err != nil {
		return errors.NewInvalidRequestError(err)
	}

	book, err := h.bookService.Get(hashId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, book)
}

func (h *BookHandler) PostBook(c echo.Context) error {
	tokenData := jwtToken.DecodeJWTToken(c.Get("user").(*jwt.Token))

	dto := &book.CreateBookDto{}
	if err := common.BindAndValidate(c, dto); err != nil {
		return errors.NewInvalidRequestError(err)
	}

	book, err := h.bookService.Create(tokenData, dto)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, book)
}

func (h *BookHandler) PutBook(c echo.Context) error {
	tokenData := jwtToken.DecodeJWTToken(c.Get("user").(*jwt.Token))

	var hashId string
	err := echo.PathParamsBinder(c).String("hash_id", &hashId).BindError()
	if err != nil {
		return errors.NewInvalidRequestError(err)
	}

	dto := &book.UpdateBookDto{}
	if err := common.BindAndValidate(c, dto); err != nil {
		return errors.NewInvalidRequestError(err)
	}

	book, err := h.bookService.Update(tokenData, hashId, dto)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, book)
}

func (h *BookHandler) DeleteBook(c echo.Context) error {
	tokenData := jwtToken.DecodeJWTToken(c.Get("user").(*jwt.Token))

	var hashId string
	err := echo.PathParamsBinder(c).String("hash_id", &hashId).BindError()
	if err != nil {
		return errors.NewInvalidRequestError(err)
	}

	err = h.bookService.Delete(tokenData, hashId)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "")
}
