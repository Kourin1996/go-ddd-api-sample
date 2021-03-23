package apierror

import (
	"github.com/labstack/echo/v4"
)

func ApiErrorHandler(err error, ctx echo.Context) {
	switch e := err.(type) {
	case *ApiError:
		ctx.JSON(e.Status, e)
		return
	}
}
