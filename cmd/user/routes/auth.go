package routes

import (
	"context"
	"fmt"

	"github.com/Surya-7890/book_store/user/config"
	"github.com/Surya-7890/book_store/user/db"
	"github.com/Surya-7890/book_store/user/gen"
	"github.com/Surya-7890/book_store/user/utils"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserService struct {
	gen.UnimplementedUserAuthServer
	DB    *gorm.DB
	Kafka *config.KafkaWriters
}

/* POST: /v1/user/login */
func (u *UserService) UserLogin(ctx context.Context, req *gen.UserLoginRequest) (*gen.UserLoginResponse, error) {
	res := &gen.UserLoginResponse{}
	username := req.GetUsername()
	password := req.GetPassword()

	if username == "" {
		res.Status = RESPONSE_FAILURE
		err_ := u.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.REQUEST_ERROR),
			Value: []byte("[user-service]: username must be provided"),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Errorf(codes.InvalidArgument, "username must be provided")
	}

	if password == "" {
		res.Status = RESPONSE_FAILURE
		err_ := u.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.REQUEST_ERROR),
			Value: []byte("[user-service]: password must be provided"),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Errorf(codes.InvalidArgument, "password must be provided")
	}

	user := &db.User{}
	if err := u.DB.Model(&db.User{}).Where("username = ?", username).First(user).Error; err != nil {
		res.Status = RESPONSE_FAILURE
		err_ := u.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.DB_ERROR),
			Value: []byte("[user-service]: " + err.Error()),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Errorf(codes.Internal, "error while logging in %s", err.Error())
	}

	if !user.IsCorrectPassword(password) {
		res.Status = RESPONSE_FAILURE
		err_ := u.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte("[user-service]: incorrect password"),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.Unauthenticated, "incorrect password")
	}
	res.Status = RESPONSE_SUCCESS
	res.User = &gen.User{
		Id:       int32(user.ID),
		Age:      user.Age,
		Username: user.Username,
		Name:     user.Name,
	}

	err_ := u.Kafka.Info.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(utils.USER_CREATE),
		Value: []byte("[user-service]: user login successful"),
	})
	if err_ != nil {
		fmt.Println(err_.Error())
	}

	return res, nil
}

/* POST: /v1/user/signup */
func (u *UserService) UserSignup(ctx context.Context, req *gen.UserSignupRequest) (*gen.UserSignupResponse, error) {
	res := &gen.UserSignupResponse{}
	username := req.GetUsername()
	name := req.GetName()
	age := req.GetAge()
	password := req.GetPassword()

	if len(username) == 0 {
		res.Status = RESPONSE_FAILURE
		err_ := u.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.REQUEST_ERROR),
			Value: []byte("[user-service]: usernane should be provided"),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.InvalidArgument, "username should be provided")
	}
	if len(name) == 0 {
		res.Status = RESPONSE_FAILURE
		err_ := u.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.REQUEST_ERROR),
			Value: []byte("[user-service]: username must be provided"),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.InvalidArgument, "name should be provided")
	}

	user := &db.User{
		Name:     name,
		Age:      age,
		Username: username,
		Password: password,
	}

	err := user.HashPassword()
	if err != nil {
		res.Status = RESPONSE_FAILURE
		err_ := u.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.DB_ERROR),
			Value: []byte("[user-service]: " + err.Error()),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.AlreadyExists, err.Error())
	}

	if user.AlreadyExists(u.DB) {
		res.Status = RESPONSE_FAILURE
		err_ := u.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.DB_ERROR),
			Value: []byte("[user-service]: username already exists"),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.AlreadyExists, "username already in use")
	}

	if err := u.DB.Create(user).Error; err != nil {
		res.Status = RESPONSE_FAILURE
		err_ := u.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.DB_ERROR),
			Value: []byte("[user-service]: " + err.Error()),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Errorf(codes.Internal, "error while creating user %s", err.Error())
	}

	res.Status = RESPONSE_SUCCESS
	res.User = &gen.User{
		Id:       int32(user.ID),
		Age:      user.Age,
		Username: user.Username,
		Name:     user.Name,
	}

	err_ := u.Kafka.Info.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(utils.USER_CREATE),
		Value: []byte("[user-service]: user created successfully"),
	})
	if err_ != nil {
		fmt.Println(err_.Error())
	}

	return res, nil
}
