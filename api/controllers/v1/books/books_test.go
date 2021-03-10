package books

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Kourin1996/go-crud-api-sample/api/models/book"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type BookServiceMock struct {
	mock.Mock
}

func (m *BookServiceMock) Create(command *book.CreateBookCommand) (*book.BookModel, error) {
	ret := m.Called(command)
	return ret.Get(0).(*book.BookModel), ret.Error(1)
}

func (m *BookServiceMock) Get(id int32) (*book.BookModel, error) {
	ret := m.Called(id)
	return ret.Get(0).(*book.BookModel), ret.Error(1)
}

func (m *BookServiceMock) Update(id int32, command *book.UpdateBookCommand) (*book.BookModel, error) {
	ret := m.Called(id, command)
	return ret.Get(0).(*book.BookModel), ret.Error(1)
}

func (m *BookServiceMock) Delete(id int32) error {
	ret := m.Called(id)
	return ret.Error(0)
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func TestPostBook(t *testing.T) {
	book := book.BookModel{
		Name:        "sensuikan1973",
		Description: "Nice",
		Price:       100,
	}
	bookJson := `{"id":0,"name":"sensuikan1973","description":"Nice","price":100}`

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	req := httptest.NewRequest(http.MethodPost, "/v1/books", strings.NewReader(bookJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	g := e.Group("/v1/books")
	mockService := &BookServiceMock{}
	mockService.On("Create", &book).Return(&book, nil)
	h := NewBookHandler(g, mockService)

	if assert.NoError(t, h.postBook(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.JSONEq(t, bookJson, rec.Body.String())
	}
}

func TestGet(t *testing.T) {
	book := book.BookModel{
		Name:        "sensuikan1973",
		Description: "Nice",
		Price:       100,
	}
	bookJson := `{"id":0,"name":"sensuikan1973","description":"Nice","price":100}`

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/books/%d", book.ID), strings.NewReader(""))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	g := e.Group("/v1/books")
	mockService := &BookServiceMock{}
	mockService.On("Get", book.ID).Return(&book, nil)
	h := NewBookHandler(g, mockService)

	if assert.NoError(t, h.Get(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, bookJson, rec.Body.String())
	}
}

func TestPutBook(t *testing.T) {
	book := book.BookModel{
		ID:          0,
		Name:        "sensuikan1973",
		Description: "Nice",
		Price:       100,
	}
	bookJson := `{"id":0,"name":"sensuikan1973","description":"Nice","price":100}`

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/v1/books/%d", book.ID), strings.NewReader(bookJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	g := e.Group("/v1/books")
	mockService := &BookServiceMock{}
	mockService.On("Update", book.ID, &book).Return(&book, nil)
	h := NewBookHandler(g, mockService)

	if assert.NoError(t, h.putBook(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, bookJson, rec.Body.String())
	}
}

func TestDelete(t *testing.T) {
	id := 0

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/books/%d", id), strings.NewReader(""))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	g := e.Group("/v1/books")
	mockService := &BookServiceMock{}
	mockService.On("Delete", id).Return(nil)
	h := NewBookHandler(g, mockService)

	if assert.NoError(t, h.Delete(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "", rec.Body.String())
	}
}
