package controllers

import (
	"github.com/go-pg/pg/v10"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Kourin1996/go-crud-api-sample/api/common"
	AuthController "github.com/Kourin1996/go-crud-api-sample/api/controllers/v1/auth"
	BooksController "github.com/Kourin1996/go-crud-api-sample/api/controllers/v1/books"
	MeController "github.com/Kourin1996/go-crud-api-sample/api/controllers/v1/me"
	BookRepository "github.com/Kourin1996/go-crud-api-sample/api/repositories/pg/book"
	UserRepository "github.com/Kourin1996/go-crud-api-sample/api/repositories/pg/user"
	AuthService "github.com/Kourin1996/go-crud-api-sample/api/services/auth"
	BookService "github.com/Kourin1996/go-crud-api-sample/api/services/book"
	UserService "github.com/Kourin1996/go-crud-api-sample/api/services/user"
)

type Config struct {
	Address string
}

func Start(config Config, db *pg.DB) error {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Validator = common.NewCustomValidator()

	g := e.Group("/v1")

	bookRepo := BookRepository.NewRepository(db)
	bookService := BookService.NewBookService(bookRepo)
	userRepo := UserRepository.NewUserRepository(db)
	userService := UserService.NewUserService(userRepo)
	authService := AuthService.NewAuthService(userService)

	BooksController.NewBookHandler(g, bookService)
	AuthController.NewAuthController(g, authService)
	MeController.NewMeController(g, userService)

	e.Logger.Fatal(e.Start(config.Address))

	return nil
}
