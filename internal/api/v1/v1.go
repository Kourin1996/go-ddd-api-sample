package v1

import (
	"net/http"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Kourin1996/go-crud-api-sample/internal/api/v1/controllers/books"
)

//todo: move to common
type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func Start() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//todo: trailing slash
	e.Validator = &CustomValidator{validator: validator.New()}

	g := e.Group("/v1")
	setRoutes(g)

	//todo: port
	e.Logger.Fatal(e.Start(":8080"))
}

func setRoutes(g *echo.Group) {
	books.SetRoutes(g)
}
