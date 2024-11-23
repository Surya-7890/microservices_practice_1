package db

import (
	"context"
	"fmt"

	"github.com/Surya-7890/book_store/admin/config"
	"github.com/Surya-7890/book_store/admin/utils"
	"github.com/segmentio/kafka-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Admin struct {
	*gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *Admin) BeforeSave(db *gorm.DB) error {
	password := a.Password
	if len(password) < 8 {
		return fmt.Errorf("error while hashing password, min length is 8")
	}

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return fmt.Errorf("error while hashing password %s", err.Error())
	}

	a.Password = string(hashedPasswordBytes)

	return nil
}

func (a *Admin) IsCorrectPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password)) == nil
}

func (a *Admin) AlreadyExists(Kafka *config.KafkaWriters, db *gorm.DB) bool {
	var count int64
	if err := db.Model(&Admin{}).Where("username = ?", a.Username).Count(&count).Error; err != nil {
		Kafka.Error.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(utils.DB_ERROR),
			Value: []byte(err.Error()),
		})
		panic(err)
	}
	return count > 0
}
