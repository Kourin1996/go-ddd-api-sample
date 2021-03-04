package books

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Kourin1996/go-crud-api-sample/pkg/models"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type BookRepositoryMock struct {
	mock.Mock
}

func (m *BookRepositoryMock) CreateBook(book *models.Book) error {
	ret := m.Called(book)
	book.ID = 0
	return ret.Error(0)
}

func (m *BookRepositoryMock) GetBook(id int) (*models.Book, error) {
	ret := m.Called(id)
	return ret.Get(0).(*models.Book), ret.Error(1)
}

func (m *BookRepositoryMock) UpdateBook(id int, book *models.Book) error {
	ret := m.Called(id, book)
	return ret.Error(0)
}

func (m *BookRepositoryMock) DeleteBook(id int) error {
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
	book := models.Book{
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
	mockRepo := &BookRepositoryMock{}
	mockRepo.On("CreateBook", &book).Return(nil)
	h := NewBookHandler(g, mockRepo)

	if assert.NoError(t, h.postBook(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.JSONEq(t, bookJson, rec.Body.String())
	}
}

func TestGetBook(t *testing.T) {
	book := models.Book{
		ID:          0,
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
	mockRepo := &BookRepositoryMock{}
	mockRepo.On("GetBook", book.ID).Return(&book, nil)
	h := NewBookHandler(g, mockRepo)

	if assert.NoError(t, h.getBook(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, bookJson, rec.Body.String())
	}
}

func TestPutBook(t *testing.T) {
	book := models.Book{
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
	mockRepo := &BookRepositoryMock{}
	mockRepo.On("UpdateBook", book.ID, &book).Return(nil)
	h := NewBookHandler(g, mockRepo)

	if assert.NoError(t, h.putBook(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, bookJson, rec.Body.String())
	}
}

func TestDeleteBook(t *testing.T) {
	id := 0

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/books/%d", id), strings.NewReader(""))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	g := e.Group("/v1/books")
	mockRepo := &BookRepositoryMock{}
	mockRepo.On("DeleteBook", id).Return(nil)
	h := NewBookHandler(g, mockRepo)

	if assert.NoError(t, h.deleteBook(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "", rec.Body.String())
	}
}