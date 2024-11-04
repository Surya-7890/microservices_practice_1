package routes

import (
	"context"
	"fmt"

	"github.com/Surya-7890/book_store/user/gen"
)

/* POST: /v1/comments/{id} */
func (u *UserService) PostComment(ctx context.Context, req *gen.PostCommentRequest) (*gen.PostCommentResponse, error) {
	res := &gen.PostCommentResponse{}
	fmt.Println("add comment")
	return res, nil
}

/* DELETE: /v1/comments/{id} */
func (u *UserService) DeleteComment(ctx context.Context, req *gen.DeleteCommentRequest) (*gen.DeleteCommentResponse, error) {
	res := &gen.DeleteCommentResponse{}
	fmt.Println("delete comment")
	return res, nil
}

/* PATCH: /v1/comments/{id} */
func (u *UserService) UpdateComment(ctx context.Context, req *gen.UpdateCommentRequest) (*gen.UpdateCommentResponse, error) {
	res := &gen.UpdateCommentResponse{}
	fmt.Println("update comment")
	return res, nil
}