package routes

import (
	"context"

	"github.com/Surya-7890/book_store/books/gen"
)

type ModifyCategoriesService struct {
	*gen.UnimplementedModifyCategoriesServer
}

func (m *ModifyCategoriesService) CreateCategory(context.Context, *gen.NewCategoryRequest) (*gen.NewCategoryResponse, error) {
	res := &gen.NewCategoryResponse{}
	return res, nil
}
func (m *ModifyCategoriesService) DeleteCategory(context.Context, *gen.DeleteCategoryRequest) (*gen.DeleteCategoryResponse, error) {
	res := &gen.DeleteCategoryResponse{}
	return res, nil
}
func (m *ModifyCategoriesService) UpdateCategory(context.Context, *gen.UpdateCategoryRequest) (*gen.UpdateCategoryResponse, error) {
	res := &gen.UpdateCategoryResponse{}
	return res, nil
}
