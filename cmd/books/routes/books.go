package routes

import (
	"context"
	"fmt"

	"github.com/Surya-7890/book_store/books/gen"
	"gorm.io/gorm"
)

type BooksService struct {
	gen.UnimplementedBooksServer
	DB *gorm.DB
}

/* GET: /v1/categories/{id}/books */
func (b *BooksService) GetBook(ctx context.Context, req *gen.GetBookRequest) (*gen.GetBookResponse, error) {
	res := &gen.GetBookResponse{}
	fmt.Println("get single book")
	return res, nil
}

/* GET: /v1/categories/{c_id}/books/{id} */
func (b *BooksService) GetBooks(ctx context.Context, req *gen.GetBooksRequest) (*gen.GetBooksResponse, error) {
	res := &gen.GetBooksResponse{}
	fmt.Println("get books")
	return res, nil
}

/* GET: /v1/categories/{id} */
func (b *BooksService) GetCategories(ctx context.Context, req *gen.GetCategoriesRequest) (*gen.GetCategoriesResponse, error) {
	res := &gen.GetCategoriesResponse{}
	fmt.Println("get categories")
	return res, nil
}
