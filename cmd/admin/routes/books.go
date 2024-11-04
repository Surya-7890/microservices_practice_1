package routes

import (
	"context"
	"fmt"

	"github.com/Surya-7890/book_store/admin/gen"
)

type AdminBooksService struct {
	gen.UnimplementedAdminBooksServer
}

/* POST: /v1/books */
func (a *AdminBooksService) NewBook(ctx context.Context, req *gen.NewBookRequest) (*gen.NewBookResponse, error) {
	res := &gen.NewBookResponse{}
	fmt.Println("new book")
	return res, nil
}

/* DELETE: /v1/books/{id} */
func (a *AdminBooksService) DeleteBooks(ctx context.Context, req *gen.DeleteBookRequest) (*gen.DeleteBookResponse, error) {
	res := &gen.DeleteBookResponse{}
	fmt.Println("delete book")
	return res, nil
}

/* PATCH: /v1/books */
func (a *AdminBooksService) UpdateBooks(ctx context.Context, req *gen.UpdateBookRequest) (*gen.UpdateBookResponse, error) {
	res := &gen.UpdateBookResponse{}
	fmt.Println("update book")
	return res, nil
}