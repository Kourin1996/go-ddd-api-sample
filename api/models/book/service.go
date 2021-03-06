package book

type IBookService interface {
	CreateBook(book *Book) (*Book, error)
	GetBook(id int) (*Book, error)
	UpdateBook(id int, book *Book) (*Book, error)
	DeleteBook(id int) error
}
