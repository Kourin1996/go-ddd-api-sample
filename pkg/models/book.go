package models

type Book struct {
	ID          int
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
}
