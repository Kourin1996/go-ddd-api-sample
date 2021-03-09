package book

type IBookService interface {
	CreateBook(book *CreateBookCommand) (*BookModel, error)
	GetBook(id int32) (*BookModel, error)
	UpdateBook(id int32, book *UpdateBookCommand) (*BookModel, error)
	DeleteBook(id int32) error
}
