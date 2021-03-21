package helper

import (
	"github.com/labstack/echo/v4"

	"github.com/Kourin1996/go-crud-api-sample/api/common"
)

func NewTestEcho() *echo.Echo {
	e := echo.New()
	e.Validator = common.NewCustomValidator()
	return e
}
