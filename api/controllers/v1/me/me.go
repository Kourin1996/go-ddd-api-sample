package auth

import (
	"net/http"

	"github.com/Kourin1996/go-crud-api-sample/api/controllers/apierror"
	"github.com/Kourin1996/go-crud-api-sample/api/controllers/middleware"
	jwtToken "github.com/Kourin1996/go-crud-api-sample/api/models/jwt"
	"github.com/Kourin1996/go-crud-api-sample/api/models/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

const BASE_PATH = "/me"

type MeController struct {
	userService user.IUserService
}

func NewMeController(group *echo.Group, userService user.IUserService) *MeController {
	c := &MeController{userService: userService}

	g := group.Group(BASE_PATH)
	g.Use(middleware.NewJwt())
	g.GET("", c.GetMe)

	return c
}

func (c *MeController) GetMe(ctx echo.Context) error {
	tokenData := jwtToken.DecodeJWTToken(ctx.Get("user").(*jwt.Token))

	me, err := c.userService.GetByHashId(tokenData.HashId)
	if err != nil {
		return apierror.NewApiError(http.StatusInternalServerError, apierror.ERROR_NOT_FOUND, err.Error())
	}

	return ctx.JSON(http.StatusOK, me)
}
