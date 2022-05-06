package model

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID                int       `gorm:"id;primaryKey;autoIncrement"`
	UserName          string    `gorm:"user_name;unique"`
	PasswordEncrypted []byte    `gorm:"password_encrypted"`
	IsAdmin           bool      `gorm:"is_admin"`
	CreateAt          time.Time `gorm:"create_at;autoCreateTime"`
	UpdateAt          time.Time `gorm:"update_at;autoUpdateTime"`
}

func GetUsers() ([]*User, error) {
	var res []*User
	result := database.Find(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return res, nil
}

func GetUserByID(ID int) (*User, error) {
	user := User{
		ID: ID,
	}
	result := database.First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUserByUserName(userName string) (*User, error) {
	user := User{}
	result := database.Where("user_name = ?", userName).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func CreateUser(user *User) error {
	result := database.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) Update() error {
	result := database.Where("id =  ?", u.ID).Select("user_name", "password_encrypted", "is_admin").Updates(u)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (u *User) Delete() error {
	result := database.Delete(u)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (u *User) SetPassword(password string) error {
	if password == "" {
		return fmt.Errorf("password is empty")
	}
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	u.PasswordEncrypted = encrypted
	return err
}

func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(u.PasswordEncrypted, []byte(password))
	return err == nil
}
