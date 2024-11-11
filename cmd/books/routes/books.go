package routes

import (
	"context"
	"strconv"

	"github.com/Surya-7890/book_store/books/db"
	"github.com/Surya-7890/book_store/books/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type BooksService struct {
	gen.UnimplementedBooksServer
	DB *gorm.DB
}

/* GET: /v1/books/{id} */
func (b *BooksService) GetBook(ctx context.Context, req *gen.GetBookRequest) (*gen.GetBookResponse, error) {
	res := &gen.GetBookResponse{}
	book := &db.Book{}
	if err := b.DB.Where("id = ?", req.Id).Find(book).Error; err != nil {
		return res, status.Errorf(codes.InvalidArgument, "error while getting book %s", err.Error())
	}
	res.Book = &gen.Book{
		Id:     strconv.Itoa(book.ID),
		Name:   book.Name,
		Price:  book.Price,
		Author: book.Author,
	}
	return res, nil
}

/* GET: /v1/books */
func (b *BooksService) GetBooks(ctx context.Context, req *gen.GetBooksRequest) (*gen.GetBooksResponse, error) {
	res := &gen.GetBooksResponse{}
	var books []db.Book
	if err := b.DB.Find(&books).Error; err != nil {
		return res, status.Errorf(codes.Internal, "error while fetching books %s", err.Error())
	}
	for _, book := range books {
		newBook := gen.Book{
			Id:     strconv.Itoa(book.ID),
			Name:   book.Name,
			Price:  book.Price,
			Author: book.Author,
		}
		res.Books = append(res.Books, &newBook)
	}
	return res, nil
}
