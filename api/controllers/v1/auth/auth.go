package auth

import (
	"net/http"

	"github.com/Kourin1996/go-crud-api-sample/api/common"
	"github.com/Kourin1996/go-crud-api-sample/api/models/auth"
	"github.com/labstack/echo/v4"
)

const BASE_PATH = "/auth"

type AuthController struct {
	authService auth.IAuthService
}

func NewAuthController(group *echo.Group, authService auth.IAuthService) *AuthController {
	c := &AuthController{authService: authService}

	g := group.Group(BASE_PATH)
	g.POST("/signup", c.SignUp)
	g.POST("/signin", c.SignIn)

	return c
}

func (c *AuthController) SignUp(ctx echo.Context) error {
	dto := &auth.SignUpDto{}
	if err := common.BindAndValidate(ctx, dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := c.authService.SignUp(dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *AuthController) SignIn(ctx echo.Context) error {
	dto := &auth.SignInDto{}
	if err := common.BindAndValidate(ctx, dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := c.authService.SignIn(dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}
