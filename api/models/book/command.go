package book

type CreateBookCommand struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
}

type UpdateBookCommand struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Price       *int64  `json:"price"`
}
