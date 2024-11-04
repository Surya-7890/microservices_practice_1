package routes

import (
	"context"
	"fmt"

	"github.com/Surya-7890/book_store/admin/gen"
)

type BookCategoryService struct {
	gen.UnimplementedBookCategoriesServer
}

/* POST: /v1/books/categories */
func (a *BookCategoryService) CreateCategory(ctx context.Context, req *gen.NewCategoryRequest) (*gen.NewCategoryResponse, error) {
	res := &gen.NewCategoryResponse{}
	fmt.Println("create category")
	return res, nil
}

/* DELETE: /v1/books/categories/{id} */
func (a *BookCategoryService) DeleteCategory(ctx context.Context, req *gen.DeleteCategoryRequest) (*gen.DeleteCategoryResponse, error) {
	res := &gen.DeleteCategoryResponse{}
	fmt.Println("delete category")
	return res, nil
}

/* PATCH: /v1/books/categories */
func (a *BookCategoryService) UpdateCategory(ctx context.Context, req *gen.UpdateCategoryRequest) (*gen.UpdateCategoryResponse, error) {
	res := &gen.UpdateCategoryResponse{}
	fmt.Println("update category")
	return res, nil
}