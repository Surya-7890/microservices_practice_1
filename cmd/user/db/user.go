package db

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Name     string `json:"name"`
	Age      int32  `json:"age"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) BeforeSave(db *gorm.DB) error {
	password := u.Password
	if len(password) < 8 {
		return fmt.Errorf("password should have a mininum length of 8")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return fmt.Errorf("error while hashing password %s", err.Error())
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) IsCorrectPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func (u *User) AlreadyExists(db *gorm.DB) bool {
	var count int64
	if err := db.Model(&User{}).Where("username = ?", u.Username).Count(&count).Error; err != nil {
		return true
	}
	return count > 0
}
