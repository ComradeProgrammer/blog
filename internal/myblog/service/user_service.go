package service

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/ComradeProgrammer/blog/internal/myblog/dal/model"
)

type UserService interface {
	ListUsers(currentUser *model.User) ([]*model.User, error)
	GetUser(currentUser *model.User, id int) (*model.User, error)
	PostUser(currentUser *model.User, newUser *model.User) (*model.User, error)
	PutUser(currentUser *model.User, id int, newUser *model.User) error
	DeleteUser(currentUser *model.User, id int) error
	ChangePassword(currentUser *model.User, oldPassword string, newPassword string) error
}

type UserServiceImpl struct {
	ServiceBase
	db *gorm.DB
}

func NewUserServiceImpl(db *gorm.DB)( *UserServiceImpl,error) {
	return &UserServiceImpl{
		ServiceBase: NewServiceBase(),
		db:          db,
	},nil
}

func (u *UserServiceImpl) ListUsers(currentUser *model.User) (res []*model.User, err error) {
	u.db.Transaction(func(tx *gorm.DB) error {
		res, err = u.userDao.GetUsers(tx)
		return err
	})
	return
}
func (u *UserServiceImpl) GetUser(currentUser *model.User, id int) (res *model.User, err error) {
	u.db.Transaction(func(tx *gorm.DB) error {
		res, err = u.userDao.GetUserByID(tx, id)
		return err
	})
	return
}
func (u *UserServiceImpl) PostUser(currentUser *model.User, newUser *model.User) (*model.User, error) {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		return u.userDao.CreateUser(tx, newUser)
	})
	return newUser, err
}

func (u *UserServiceImpl) DeleteUser(currentUser *model.User, id int) error {
	if currentUser == nil || !currentUser.IsAdmin {
		return fmt.Errorf("only admin user can delete users")
	}
	return u.db.Transaction(func(tx *gorm.DB) error {
		return u.userDao.Delete(tx, &model.User{ID: id})
	})
}
func (u *UserServiceImpl) PutUser(currentUser *model.User, id int, newUser *model.User) error {
	if currentUser == nil || currentUser.ID != id && !currentUser.IsAdmin {
		return fmt.Errorf("only admin user or user itself can delete a user")
	}
	newUser.ID = id
	return u.db.Transaction(func(tx *gorm.DB) error {
		return u.userDao.Update(tx, newUser)
	})
}


func (u *UserServiceImpl) ChangePassword(currentUser *model.User, oldPassword string, newPassword string) error {
	if currentUser == nil {
		return fmt.Errorf("changing password requires logging in")
	}

	if ok := currentUser.VerifyPassword(oldPassword); !ok {
		return fmt.Errorf("incorrect old password")
	}
	if err := currentUser.SetPassword(newPassword); err != nil {
		return fmt.Errorf("failed to set new password: %v", err)
	}
	return u.db.Transaction(func(tx *gorm.DB) error {
		return u.userDao.Update(tx, currentUser)
	})

}
