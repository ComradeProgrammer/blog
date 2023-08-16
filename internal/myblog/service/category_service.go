package service

import (
	"fmt"

	"github.com/ComradeProgrammer/blog/internal/myblog/dal/model"
	"gorm.io/gorm"
)

type CategoryService interface {
	ListCategories(currentUser *model.User) ([]*model.Category, error)
	GetCategory(currentUser *model.User, id int) (*model.Category, error)
	PostCategory(currentUser *model.User, category *model.Category) error
	PutCategory(currentUser *model.User, id int, category *model.Category) error
	DeleteCategory(currentUser *model.User, id int) error
}

type CategoryServiceImpl struct {
	ServiceBase

	db *gorm.DB
}

func NewCategoryServiceImpl(db *gorm.DB) (*CategoryServiceImpl, error) {
	return &CategoryServiceImpl{
		ServiceBase: NewServiceBase(),
		db:          db,
	}, nil
}
func (c *CategoryServiceImpl) ListCategories(currentUser *model.User) (res []*model.Category, err error) {
	c.db.Transaction(func(tx *gorm.DB) error {
		res, err = c.categoryDao.GetCategories(tx)
		return err
	})
	return
}
func (c *CategoryServiceImpl) GetCategory(currentUser *model.User, id int) (res *model.Category, err error) {
	c.db.Transaction(func(tx *gorm.DB) error {
		res, err = c.categoryDao.GetCategory(tx, id)
		return err
	})
	return

}
func (c *CategoryServiceImpl) PostCategory(currentUser *model.User, category *model.Category) error {
	if currentUser == nil || !currentUser.IsAdmin {
		return fmt.Errorf("only admin can post category")
	}
	if category.Name == "" {
		return fmt.Errorf("name field must not be empty")
	}
	return c.db.Transaction(func(tx *gorm.DB) error {
		return c.categoryDao.CreateCategory(tx, category)
	})
}
func (c *CategoryServiceImpl) PutCategory(currentUser *model.User, id int, category *model.Category) error {
	if currentUser == nil || !currentUser.IsAdmin {
		return fmt.Errorf("only admin can put category")
	}

	category.ID = id
	if category.Name == "" {
		return fmt.Errorf("name field must not be empty")
	}

	return c.db.Transaction(func(tx *gorm.DB) error {
		return c.categoryDao.Update(tx, category)
	})
}
func (c *CategoryServiceImpl) DeleteCategory(currentUser *model.User, id int) error {
	if currentUser == nil || !currentUser.IsAdmin {
		return fmt.Errorf("only admin can delete log")
	}
	return c.db.Transaction(func(tx *gorm.DB) error {
		return c.categoryDao.Delete(tx, &model.Category{
			ID: id,
		})
	})
}
