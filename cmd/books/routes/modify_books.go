package routes

import (
	"context"

	"github.com/Surya-7890/book_store/books/gen"
	"gorm.io/gorm"
)

type ModifyBooksService struct {
	gen.UnimplementedModifyBooksServer
	DB *gorm.DB
}

func (m *ModifyBooksService) NewBook(ctx context.Context, req *gen.NewBookRequest) (*gen.NewBookResponse, error) {
	res := &gen.NewBookResponse{}
	return res, nil
}
func (m *ModifyBooksService) DeleteBooks(ctx context.Context, req *gen.DeleteBookRequest) (*gen.DeleteBookResponse, error) {
	res := &gen.DeleteBookResponse{}
	return res, nil
}
func (m *ModifyBooksService) UpdateBooks(ctx context.Context, req *gen.UpdateBookRequest) (*gen.UpdateBookResponse, error) {
	res := &gen.UpdateBookResponse{}
	return res, nil
}
