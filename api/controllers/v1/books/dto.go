package books

type PostBookDto struct {
	Name        string `validate:"required,min=1,max=255"`
	Description string `validate:"required,max=1024"`
	Price       int64  `validate:"required"`
}

type PutBookDto struct {
	Name        *string `validate:"min=1,max=255"`
	Description *string `validate:"max=1024"`
	Price       *int64  `validate:""`
}

type DeleteBookDto struct {
	Id string `path:"id"`
}
