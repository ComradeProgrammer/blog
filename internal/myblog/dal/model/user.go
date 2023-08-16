package model

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int       `gorm:"id;primaryKey;autoIncrement"`
	UserName          string    `gorm:"user_name;unique" json:"userName"`
	PasswordEncrypted []byte    `gorm:"password_encrypted" json:"-"` //used to store the password in database
	Password          string    `gorm:"-" json:"password"`           //use to pass the password in body
	IsAdmin           bool      `gorm:"is_admin" json:"isAdmin"`
	CreateAt          time.Time `gorm:"create_at;autoCreateTime" json:"createAt"`
	UpdateAt          time.Time `gorm:"update_at;autoUpdateTime" json:"updateAt"`
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
