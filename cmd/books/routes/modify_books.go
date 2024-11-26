package routes

import (
	"context"
	"fmt"
	"strings"

	"github.com/Surya-7890/book_store/books/config"
	"github.com/Surya-7890/book_store/books/db"
	"github.com/Surya-7890/book_store/books/gen"
	"github.com/Surya-7890/book_store/books/utils"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ModifyBooksService struct {
	gen.UnimplementedModifyBooksServer
	DB    *gorm.DB
	Kafka *config.KafkaWriters
}

// post: /v1/books
func (m *ModifyBooksService) NewBook(ctx context.Context, req *gen.NewBookRequest) (*gen.NewBookResponse, error) {
	res := &gen.NewBookResponse{}
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		res.Status = RESPONSE_FAILURE
		err_ := m.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte("[books-service]: invalid header"),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.InvalidArgument, "invalid header")
	}

	errors := md.Get("auth-error")
	if len(errors) != 0 {
		err_ := m.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte("[books-service]: " + strings.Join(errors, " ")),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.PermissionDenied, strings.Join(errors, ", "))
	}

	role := md.Get("role")
	if role[0] != "admin" {
		res.Status = RESPONSE_FAILURE
		err_ := m.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte("[books-service]: operation not permitted for the user"),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.Unauthenticated, "operation not permitted")
	}

	name := req.GetName()
	author := req.GetAuthor()
	price := req.GetPrice()

	if len(name) == 0 {
		res.Status = RESPONSE_FAILURE
		err_ := m.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte("[books-service]: book name is required for creation"),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.InvalidArgument, "book name is required")
	}

	if len(author) == 0 {
		res.Status = RESPONSE_FAILURE
		err_ := m.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte("[books-service]: author name is required for creation"),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.InvalidArgument, "author name is required")
	}

	book := &db.Book{
		Name:   name,
		Author: author,
		Price:  price,
	}
	tx := m.DB.Create(book)

	if tx.Error != nil {
		res.Status = RESPONSE_FAILURE
		err_ := m.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte("[books-service]: " + tx.Error.Error()),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Errorf(codes.Internal, "error while creating book %s", tx.Error.Error())
	}

	if tx.RowsAffected == 0 {
		res.Status = RESPONSE_FAILURE
		err_ := m.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte("[books-service]: error while creating book"),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.Internal, "error while creating book")
	}
	res.Status = RESPONSE_SUCCESS
	return res, nil
}

// delete: /v1/books/{id}
func (m *ModifyBooksService) DeleteBooks(ctx context.Context, req *gen.DeleteBookRequest) (*gen.DeleteBookResponse, error) {
	res := &gen.DeleteBookResponse{}
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		res.Status = RESPONSE_FAILURE
		err_ := m.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte("[books-service]: invalid header"),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.InvalidArgument, "invalid header")
	}

	errors := md.Get("auth-error")
	if len(errors) != 0 {
		err_ := m.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte("[books-service]: " + strings.Join(errors, " ")),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.PermissionDenied, strings.Join(errors, ", "))
	}

	role := md.Get("role")
	if role[0] != "admin" {
		res.Status = RESPONSE_FAILURE
		err_ := m.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte("[books-service]: operation not permitted"),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.Unauthenticated, "operation not permitted")
	}

	if err := m.DB.Where("id = ?", req.Id).Delete(&db.Book{}).Error; err != nil {
		res.Status = RESPONSE_FAILURE
		err_ := m.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte("[books-service]: " + err.Error()),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Errorf(codes.Internal, "error while deleting book %s", err.Error())
	}
	res.Status = RESPONSE_SUCCESS
	return res, nil
}

// patch: /v1/books
func (m *ModifyBooksService) UpdateBooks(ctx context.Context, req *gen.UpdateBookRequest) (*gen.UpdateBookResponse, error) {
	res := &gen.UpdateBookResponse{}
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		res.Status = RESPONSE_FAILURE
		err_ := m.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte("[books-service]: invalid header"),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.InvalidArgument, "invalid header")
	}

	errors := md.Get("auth-error")
	if len(errors) != 0 {
		err_ := m.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte("[books-service]: " + strings.Join(errors, " ")),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.PermissionDenied, strings.Join(errors, ", "))
	}

	role := md.Get("role")
	fmt.Println(role)
	if role[0] != "admin" {
		res.Status = RESPONSE_FAILURE
		err_ := m.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte("[books-service]: operation not permitted"),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.Unauthenticated, "operation not permitted")
	}

	updatedBook := &db.Book{}
	author := req.GetAuthor()
	name := req.GetName()
	price := req.GetPrice()
	id := req.GetId()
	if len(author) != 0 {
		updatedBook.Author = author
	}
	if len(name) != 0 {
		updatedBook.Name = name
	}
	if price != 0.00 {
		updatedBook.Price = price
	}

	if err := m.DB.Where("id = ?", id).Updates(updatedBook).Error; err != nil {
		res.Status = RESPONSE_FAILURE
		err_ := m.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte("[books-service]: " + err.Error()),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Errorf(codes.Internal, "error while updating book %s", err.Error())
	}
	res.Status = RESPONSE_SUCCESS
	return res, nil
}
