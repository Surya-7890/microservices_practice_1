package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Surya-7890/book_store/user/config"
	"github.com/Surya-7890/book_store/user/db"
	"github.com/Surya-7890/book_store/user/gen"
	"github.com/Surya-7890/book_store/user/utils"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserProfileService struct {
	gen.UnimplementedUserProfileServer
	DB    *gorm.DB
	Kafka *config.KafkaWriters
}

/* GET: /v1/user/profile */
func (u *UserProfileService) GetUser(ctx context.Context, req *gen.GetUserRequest) (*gen.GetUserResponse, error) {
	res := &gen.GetUserResponse{}
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return res, status.Error(codes.InvalidArgument, "invalid header")
	}

	errors := md.Get("auth-error")
	if len(errors) != 0 {
		err_ := u.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte(strings.Join(errors, " ")),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.PermissionDenied, strings.Join(errors, ", "))
	}

	user_string := md.Get("user")[0]
	if user_string == "" {
		err_ := u.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte("invalid user"),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.Unauthenticated, "invalid user")
	}

	var user db.User

	err := json.Unmarshal([]byte(user_string), &user)
	if err != nil {
		err_ := u.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.INTERNAL_ERROR),
			Value: []byte(err.Error()),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.Unauthenticated, "invalid user found")
	}

	var user_ db.User
	if err = u.DB.First(&user_, user.ID).Error; err != nil {
		err_ := u.Kafka.Info.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.USER_INFO),
			Value: []byte(fmt.Sprintf("[user-service]: " + err.Error())),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.Internal, err.Error())
	}
	res.Age = user_.Age
	res.Name = user_.Name
	res.Username = user_.Username

	err_ := u.Kafka.Info.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(utils.USER_INFO),
		Value: []byte(fmt.Sprintf("user profile requested %v", user.ID)),
	})
	if err_ != nil {
		fmt.Println(err_.Error())
	}
	return res, nil
}

/* DELETE: /v1/user/profile */
func (u *UserProfileService) DeleteUser(ctx context.Context, req *gen.DeleteUserRequest) (*gen.DeleteUserResponse, error) {
	res := &gen.DeleteUserResponse{}
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		err_ := u.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte("invalid header"),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.InvalidArgument, "invalid header")
	}

	errors := md.Get("auth-error")
	if len(errors) != 0 {
		err_ := u.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte(strings.Join(errors, " ")),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.PermissionDenied, strings.Join(errors, ", "))
	}

	user_string := md.Get("user")[0]
	if user_string == "" {
		err_ := u.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte("invalid user"),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.Unauthenticated, "invalid user")
	}

	var user db.User

	err := json.Unmarshal([]byte(user_string), &user)
	if err != nil {
		err_ := u.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.INTERNAL_ERROR),
			Value: []byte(err.Error()),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Errorf(codes.Unauthenticated, "invalid user found %s", err.Error())
	}

	if user.Username == "" {
		err_ := u.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.INTERNAL_ERROR),
			Value: []byte("user not found"),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.Unauthenticated, "error with the user found inside token")
	}

	if err := u.DB.Where("username = ?", user.Username).Delete(&db.User{}).Error; err != nil {
		res.Status = RESPONSE_FAILURE
		err_ := u.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.DB_ERROR),
			Value: []byte(err.Error()),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Errorf(codes.Internal, "error while deleting user %s", err.Error())
	}

	res.Status = RESPONSE_SUCCESS
	err_ := u.Kafka.Info.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(utils.USER_INFO),
		Value: []byte("deleted user successfully"),
	})
	if err_ != nil {
		fmt.Println(err_.Error())
	}
	return res, nil
}

/* PATCH: /v1/user/profile */
func (u *UserProfileService) UpdateUser(ctx context.Context, req *gen.UpdateUserRequest) (*gen.UpdateUserResponse, error) {
	res := &gen.UpdateUserResponse{}
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		err_ := u.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte("[user-service]: invalid header"),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.InvalidArgument, "invalid header")
	}

	errors := md.Get("auth-error")
	if len(errors) != 0 {
		err_ := u.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte("[user-service]: " + strings.Join(errors, " ")),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.PermissionDenied, strings.Join(errors, ", "))
	}

	user_string := md.Get("user")[0]
	if user_string == "" {
		err_ := u.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.INTERNAL_ERROR),
			Value: []byte("[user-service]: invalid user"),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.Unauthenticated, "invalid user")
	}

	var user db.User

	err := json.Unmarshal([]byte(user_string), &user)
	if err != nil {
		err_ := u.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.INTERNAL_ERROR),
			Value: []byte("[user-service]: invalid user found"),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Error(codes.Unauthenticated, "invalid user found")
	}

	username := req.GetUsername()
	name := req.GetName()
	age := req.GetAge()

	updates := &db.User{}

	if len(username) > 0 {
		updates.Username = username
	}

	if len(name) > 0 {
		updates.Name = name
	}

	if age != 0 {
		updates.Age = age
	}

	updates.Password = user.Password

	if err := u.DB.Where("username = ?", user.Username).Updates(updates).Error; err != nil {
		res.Status = RESPONSE_FAILURE
		err_ := u.Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.DB_ERROR),
			Value: []byte("[user-service]: operation not permitted"),
		})
		if err_ != nil {
			fmt.Println(err_.Error())
		}
		return res, status.Errorf(codes.Internal, "error while updating user %s", err.Error())
	}

	res.Status = RESPONSE_SUCCESS

	return res, nil
}
