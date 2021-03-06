package controllers

import (
	"github.com/go-pg/pg/v10"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Kourin1996/go-crud-api-sample/api/common"
	"github.com/Kourin1996/go-crud-api-sample/api/controllers/v1/books"
	BookRepository "github.com/Kourin1996/go-crud-api-sample/api/repositories/pg/book"
	BookService "github.com/Kourin1996/go-crud-api-sample/api/services/book"
)

type Config struct {
	Address string
}

func Start(config Config, db *pg.DB) error {
	e := echo.New()
	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Validator = common.NewCustomValidator()

	g := e.Group("/v1")

	bookRepo := BookRepository.NewRepository(db)
	bookService := BookService.NewBookService(bookRepo)
	books.NewBookHandler(g, bookService)

	e.Logger.Fatal(e.Start(config.Address))

	return nil
}
