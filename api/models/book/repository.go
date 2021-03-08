package book

type IBookRepository interface {
	CreateBook(book *CreateBookCommand) (*BookResult, error)
	GetBook(id int) (*BookResult, error)
	UpdateBook(id int, newData *UpdateBookCommand) (*BookResult, error)
	DeleteBook(id int) error
}
