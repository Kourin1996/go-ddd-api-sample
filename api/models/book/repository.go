package book

type IBookRepository interface {
	Create(*Book) (*Book, error)
	Get(int64) (*Book, error)
	Update(int64, *UpdateBook) (*Book, error)
	Delete(int64) error
}
