package middleware

import (
	"github.com/Kourin1996/go-crud-api-sample/api/constants"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewJwt() echo.MiddlewareFunc {
	return middleware.JWT([]byte(constants.JWT_SECRET))
}
