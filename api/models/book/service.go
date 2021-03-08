package book

type IBookService interface {
	CreateBook(book *CreateBookCommand) (*BookResult, error)
	GetBook(id int) (*BookResult, error)
	UpdateBook(id int, book *UpdateBookCommand) (*BookResult, error)
	DeleteBook(id int) error
}
