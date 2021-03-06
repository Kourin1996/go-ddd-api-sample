package v1

import (
	"net/http"
	"os"

	"github.com/go-pg/pg/v10"
	validator "github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Kourin1996/go-crud-api-sample/api/controllers/v1/books"
	BookRepository "github.com/Kourin1996/go-crud-api-sample/api/repositories/book"
	BookService "github.com/Kourin1996/go-crud-api-sample/api/services/book"
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

type Config struct {
	Address    string
	DbAddress  string
	DbUser     string
	DbPassword string
	DbDatabase string
}

func LoadEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func LoadConfig() (Config, error) {
	var config Config

	err := godotenv.Load()
	if err != nil {
		return config, err
	}

	config.Address = LoadEnv("API_ADDRESS", ":8080")
	config.DbAddress = LoadEnv("DB_ADDRESS", ":5432")
	config.DbUser = LoadEnv("DB_USER", "")
	config.DbPassword = LoadEnv("DB_PASSWORD", "")
	config.DbDatabase = LoadEnv("DB_DATABASE", "")

	return config, err
}

func Start() error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//todo: trailing slash
	e.Validator = &CustomValidator{validator: validator.New()}

	g := e.Group("/v1")

	db := pg.Connect(&pg.Options{
		Addr:     config.DbAddress,
		User:     config.DbUser,
		Password: config.DbPassword,
		Database: config.DbDatabase,
	})
	bookRepo := BookRepository.NewRepository(db)
	bookService := BookService.NewBookService(bookRepo)
	books.NewBookHandler(g, bookService)

	e.Logger.Fatal(e.Start(config.Address))

	return nil
}
