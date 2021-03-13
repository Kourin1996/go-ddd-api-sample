package book

type IBookService interface {
	Create(book *CreateBookDto) (*Book, error)
	Get(hashId string) (*Book, error)
	Update(hashId string, book *UpdateBookDto) (*Book, error)
	Delete(hashId string) error
}
