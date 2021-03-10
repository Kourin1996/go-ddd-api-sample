package book

type IBookRepository interface {
	Create(*CreateBookCommand) (*BookModel, error)
	Get(int32) (*BookModel, error)
	Update(int32, *UpdateBookCommand) (*BookModel, error)
	Delete(int32) error
}
