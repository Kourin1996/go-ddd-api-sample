package book

type IBookRepository interface {
	CreateBook(book *Book) (*Book, error)
	GetBook(id int) (*Book, error)
	UpdateBook(id int, newData *Book) (*Book, error)
	DeleteBook(id int) error
}
