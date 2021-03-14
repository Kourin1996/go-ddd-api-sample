package common

import (
	"github.com/labstack/echo/v4"
)

func BindAndValidate(ctx echo.Context, x interface{}) error {
	if err := ctx.Bind(x); err != nil {
		return err
	}
	if err := ctx.Validate(x); err != nil {
		return err
	}
	return nil
}
