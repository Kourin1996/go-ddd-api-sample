package book

type IBookRepository interface {
	CreateBook(*CreateBookCommand) (*BookModel, error)
	GetBook(int32) (*BookModel, error)
	UpdateBook(int32, *UpdateBookCommand) (*BookModel, error)
	DeleteBook(int32) error
}
