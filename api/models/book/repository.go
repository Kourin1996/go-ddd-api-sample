package book

type IBookRepository interface {
	CreateBook(*CreateBookCommand) (*BookResult, error)
	GetBook(int32) (*BookResult, error)
	UpdateBook(int32, *UpdateBookCommand) (*BookResult, error)
	DeleteBook(int32) error
}
