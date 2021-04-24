package book_test

import (
	"fmt"
	"strings"
	"testing"
	"time"
	
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	type Args struct {
		hashId string
		book   *book.Book
		err    error
	}
	type Expected struct {
		isRepoCalled bool
		id           int64
		err          error
	}

	timestamp := time.Now()
	tests := []struct {
		name     string
		args     Args
		expected Expected
	}{
		{
			name: "should be failed because hash id is wrong",
			args: Args{
				hashId: helper.EncodeID("Test", 1),
				book:   models.NewFakeBookWithID(1, "Test Book", "This is test book", 1000, timestamp, timestamp),
				err:    nil,
			},
			expected: Expected{
				isRepoCalled: false,
				id:           1,
				err:          errors.NewInvalidDataError(fmt.Errorf("mismatch between encode and decode: %s start  re-encoded.", helper.EncodeID("Test", 1))),
			},
		},
		{
			name: "should be failed because repository returns error",
			args: Args{
				hashId: helper.EncodeID("Book", 1),
				book:   models.NewFakeBookWithID(1, "Test Book", "This is test book", 1000, timestamp, timestamp),
				err:    fmt.Errorf("Query Error"),
			},
			expected: Expected{
				isRepoCalled: true,
				id:           1,
				err:          fmt.Errorf("Query Error"),
			},
		},
		{
			name: "should be failed because repository returns nil",
			args: Args{
				hashId: helper.EncodeID("Book", 1),
				book:   nil,
				err:    nil,
			},
			expected: Expected{
				isRepoCalled: true,
				id:           1,
				err:          errors.NewNotFoundError(fmt.Errorf("Book not found")),
			},
		},
		{
			name: "should be success",
			args: Args{
				hashId: helper.EncodeID("Book", 1),
				book:   models.NewFakeBookWithID(1, "Test Book", "This is test book", 1000, time.Unix(0, 0), time.Unix(0, 0)),
				err:    nil,
			},
			expected: Expected{
				isRepoCalled: true,
				id:           1,
				err:          nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockBookRepo := &repositories.MockBookRepository{}
			mockBookRepo.On("Get", tt.expected.id).Return(tt.args.book, tt.args.err)
			bookService := BookService.NewService(mockBookRepo)

			res, err := bookService.Get(tt.args.hashId)
			if tt.expected.isRepoCalled {
				mockBookRepo.AssertExpectations(t)
			} else {
				mockBookRepo.AssertNotCalled(t, "Get", tt.expected.id)
			}
			if tt.expected.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.args.book, res)
			} else {
				assert.Error(t, err)
				assert.IsType(t, tt.expected.err, err)
				assert.Contains(t, err.Error(), tt.expected.err.Error())
			}
		})
	}
}

func TestCreate(t *testing.T) {
	type Args struct {
		dto  *book.CreateBookDto
		book *book.Book
		err  error
	}
	type Expected struct {
		isRepoCalled bool
		err          error
	}

	tests := []struct {
		name     string
		args     Args
		expected Expected
	}{
		{
			name: "should be failed because length of name is less than 1",
			args: Args{
				dto: &book.CreateBookDto{
					Name:        "",
					Description: "This is test book",
					Price:       1000,
				},
				book: nil,
				err:  nil,
			},
			expected: Expected{
				isRepoCalled: false,
				err:          errors.NewInvalidDataError(fmt.Errorf("name must be longer than 0 and shorter than 257")),
			},
		},
		{
			name: "should be failed because length of name is greater than 256",
			args: Args{
				dto: &book.CreateBookDto{
					Name:        strings.Repeat("a", 300),
					Description: "This is test book",
					Price:       1000,
				},
				book: nil,
				err:  nil,
			},
			expected: Expected{
				isRepoCalled: false,
				err:          errors.NewInvalidDataError(fmt.Errorf("name must be longer than 0 and shorter than 257")),
			},
		},
		{
			name: "should be failed because length of description is greater than 1024",
			args: Args{
				dto: &book.CreateBookDto{
					Name:        "Test Book",
					Description: strings.Repeat("a", 1030),
					Price:       1000,
				},
				book: nil,
				err:  nil,
			},
			expected: Expected{
				isRepoCalled: false,
				err:          errors.NewInvalidDataError(fmt.Errorf("description must be shorter than 1025")),
			},
		},
		{
			name: "should be failed because repository throw error",
			args: Args{
				dto: &book.CreateBookDto{
					Name:        "Test Book",
					Description: "This is test book",
					Price:       1000,
				},
				book: func() *book.Book {
					b, _ := book.NewBook(
						&book.CreateBookDto{
							Name:        "Test Book",
							Description: "This is test book",
							Price:       1000,
						},
					)
					return b
				}(),
				err: fmt.Errorf("failed to execute query"),
			},
			expected: Expected{
				isRepoCalled: true,
				err:          fmt.Errorf("failed to execute query"),
			},
		},
		{
			name: "should be success",
			args: Args{
				dto: &book.CreateBookDto{
					Name:        "Test Book",
					Description: "This is test book",
					Price:       1000,
				},
				book: func() *book.Book {
					b, _ := book.NewBook(
						&book.CreateBookDto{
							Name:        "Test Book",
							Description: "This is test book",
							Price:       1000,
						},
					)
					return b
				}(),
				err: nil,
			},
			expected: Expected{
				isRepoCalled: true,
				err:          nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockBookRepo := &repositories.MockBookRepository{}
			mockBookRepo.On("Create", tt.args.book).Return(tt.args.book, tt.args.err)
			bookService := BookService.NewService(mockBookRepo)

			res, err := bookService.Create(tt.args.dto)
			if tt.expected.isRepoCalled {
				mockBookRepo.AssertExpectations(t)
			} else {
				mockBookRepo.AssertNotCalled(t, "Create", tt.args.book)
			}
			if tt.expected.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.args.book, res)
			} else {
				assert.Error(t, err)
				assert.IsType(t, tt.expected.err, err)
				assert.Contains(t, err.Error(), tt.expected.err.Error())
			}
		})
	}

}
