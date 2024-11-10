package db

import (
	"fmt"

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

func (a *Admin) AlreadyExists(db *gorm.DB) bool {
	var count int64
	if err := db.Model(&Admin{}).Where("username = ?", a.Username).Count(&count).Error; err != nil {
		panic(err)
	}
	return count > 0
}
