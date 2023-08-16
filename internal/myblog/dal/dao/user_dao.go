package dao

import (
	"gorm.io/gorm"

	"github.com/ComradeProgrammer/blog/internal/myblog/dal/model"
)

type UserDao interface {
	GetUsers(database *gorm.DB) ([]*model.User, error)
	GetUserByID(database *gorm.DB, ID int) (*model.User, error)
	GetUserByUserName(database *gorm.DB, userName string) (*model.User, error)
	CreateUser(database *gorm.DB, user *model.User) error
	Update(database *gorm.DB, u *model.User) error
	Delete(database *gorm.DB, u *model.User) error
}

type UserDaoImpl struct {
}

func (*UserDaoImpl) GetUsers(database *gorm.DB) ([]*model.User, error) {
	var res []*model.User
	result := database.Find(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return res, nil
}
func (*UserDaoImpl) GetUserByID(database *gorm.DB, ID int) (*model.User, error) {
	user := model.User{
		ID: ID,
	}
	result := database.First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (*UserDaoImpl) GetUserByUserName(database *gorm.DB, userName string) (*model.User, error) {
	user := model.User{}
	result := database.Where("user_name = ?", userName).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil

}
func (*UserDaoImpl) CreateUser(database *gorm.DB, user *model.User) error {
	result := database.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (*UserDaoImpl) Update(database *gorm.DB, u *model.User) error {
	result := database.Where("id =  ?", u.ID).Select("user_name", "password_encrypted", "is_admin").Updates(u)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
func (*UserDaoImpl) Delete(database *gorm.DB, u *model.User) error {
	result := database.Delete(u)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
