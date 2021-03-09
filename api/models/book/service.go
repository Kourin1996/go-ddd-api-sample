package book

type IBookService interface {
	CreateBook(book *CreateBookCommand) (*BookResult, error)
	GetBook(id int32) (*BookResult, error)
	UpdateBook(id int32, book *UpdateBookCommand) (*BookResult, error)
	DeleteBook(id int32) error
}
