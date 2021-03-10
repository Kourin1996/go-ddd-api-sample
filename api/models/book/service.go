package book

type IBookService interface {
	Create(book *CreateBookCommand) (*BookModel, error)
	Get(id int32) (*BookModel, error)
	Update(id int32, book *UpdateBookCommand) (*BookModel, error)
	Delete(id int32) error
}
