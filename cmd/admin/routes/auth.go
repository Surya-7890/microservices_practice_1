package routes

import (
	"context"
	"fmt"

	"github.com/Surya-7890/book_store/admin/config"
	"github.com/Surya-7890/book_store/admin/db"
	"github.com/Surya-7890/book_store/admin/gen"
	"github.com/Surya-7890/book_store/admin/utils"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type AdminAuthService struct {
	gen.UnimplementedAdminAuthServer
	DB    *gorm.DB
	Kafka *config.KafkaWriters
}

/* POST: /v1/admin/login */
func (a *AdminAuthService) AdminLogin(ctx context.Context, req *gen.AdminLoginRequest) (*gen.AdminLoginResponse, error) {
	res := &gen.AdminLoginResponse{}
	admin := db.Admin{}

	a.DB.Where("username = ?", req.GetUsername()).First(&admin)

	if !admin.IsCorrectPassword(req.GetPassword()) {
		a.Kafka.Error.WriteMessages(ctx, kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte("[admin-service]: incorrect password"),
		})
		res.Status = RESPONSE_FAILURE
		return res, status.Error(codes.PermissionDenied, "Incorrect Password")
	}

	res.Status = RESPONSE_SUCCESS
	res.User = &gen.Admin{
		Id:       int32(admin.ID),
		Username: admin.Username,
	}
	a.Kafka.Info.WriteMessages(ctx, kafka.Message{
		Key:   []byte(utils.ADMIN_LOGIN),
		Value: []byte(fmt.Sprintf("[admin-service]: admin login successful: (%d)", admin.ID)),
	})
	return res, nil
}

/* POST: /v1/admin/create */
func (a *AdminAuthService) AdminCreate(ctx context.Context, req *gen.AdminCreateRequest) (*gen.AdminCreateResponse, error) {
	res := &gen.AdminCreateResponse{}

	admin := db.Admin{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	}
	if admin.AlreadyExists(a.Kafka, a.DB) {
		a.Kafka.Error.WriteMessages(ctx, kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte("[admin-service]: admin already exists"),
		})
		res.Status = RESPONSE_FAILURE
		return res, status.Error(codes.AlreadyExists, "Admin With The Username Already Exists")
	}
	tx := a.DB.Create(&admin)

	// handle errors while creating
	if tx.Error != nil {
		res.Status = RESPONSE_FAILURE
		a.Kafka.Error.WriteMessages(ctx, kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte("[admin-service]: " + tx.Error.Error()),
		})
		return res, status.Errorf(codes.Internal, "Unable To Create Admin: %s", tx.Error.Error())
	}

	// handle rows affected on creation
	if tx.RowsAffected == 0 {
		res.Status = RESPONSE_FAILURE
		a.Kafka.Error.WriteMessages(ctx, kafka.Message{
			Key:   []byte(utils.AUTH_ERROR),
			Value: []byte("[admin-service]: unalbe to create admin"),
		})
		return res, status.Error(codes.Unknown, "Unable To Create Admin")
	}

	res.Status = RESPONSE_SUCCESS
	res.User = &gen.Admin{
		Id:       int32(admin.ID),
		Username: admin.Username,
	}

	a.Kafka.Info.WriteMessages(ctx, kafka.Message{
		Key:   []byte(utils.ADMIN_CREATE),
		Value: []byte(fmt.Sprintf("[admin-service]: admin created successfully: (%d)", admin.ID)),
	})

	return res, nil
}
