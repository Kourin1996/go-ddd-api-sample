package auth

import (
	"net/http"

	"github.com/Kourin1996/go-crud-api-sample/api/controllers/middleware"
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
	userToken := ctx.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	hashId := claims["hash_id"].(string)

	me, err := c.userService.GetByHashId(hashId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, me)
}
