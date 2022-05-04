package model

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int    `gorm:"id;primaryKey;autoIncrement"`
	UserName          string `gorm:"user_name;unique"`
	PasswordEncrypted []byte `gorm:"password_encrypted"`
	IsAdmin           bool   `gorm:"is_admin"`
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
	result := database.Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUserByUserName(userName string) (*User, error) {
	user := User{
		UserName: userName,
	}
	result := database.Find(&user)
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
	result := database.Save(u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) Delete() error {
	result := database.Delete(u)
	if result.Error != nil {
		return result.Error
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
