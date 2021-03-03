package books

import (
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

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

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

func TestPostBook(t *testing.T) {
	book := models.Book{
		Name:        "sensuikan1973",
		Description: "Nice",
		Price:       100,
	}
	bookJson := `{"Name":"sensuikan1973", "Description": "Nice", "Price":100}`

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
		t.Errorf("Response body: %s\n", rec.Body.String())
	}
}
