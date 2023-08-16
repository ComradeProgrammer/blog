package service

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/ComradeProgrammer/blog/internal/myblog/dal/model"
)

type BlogService interface {
	ListBlogs(currentUser *model.User) ([]*model.Blog, error)
	GetBlog(currentUser *model.User, id int) (*model.Blog, error)
	PostBlog(currentUser *model.User, blog *model.Blog) error
	PutBlog(currentUser *model.User, id int, blog *model.Blog) error
	DeleteBlog(currentUser *model.User, id int) error
}

type BlogServiceImpl struct {
	ServiceBase

	db *gorm.DB
}

func NewBlogServiceImpl(db *gorm.DB) (*BlogServiceImpl, error) {
	return &BlogServiceImpl{
		ServiceBase: NewServiceBase(),
		db:          db,
	}, nil
}

func (b *BlogServiceImpl) ListBlogs(currentUser *model.User) (res []*model.Blog, err error) {
	b.db.Transaction(func(tx *gorm.DB) error {
		res, err = b.blogDao.GetBlogs(tx)
		return err
	})
	return
}

func (b *BlogServiceImpl) GetBlog(currentUser *model.User, id int) (res *model.Blog, err error) {
	b.db.Transaction(func(tx *gorm.DB) error {
		res, err = b.blogDao.GetBlog(tx, id)
		return err
	})
	return
}

func (b *BlogServiceImpl) PostBlog(currentUser *model.User, blog *model.Blog) error {
	if currentUser == nil || !currentUser.IsAdmin {
		return fmt.Errorf("only admin can post log")
	}

	return b.db.Transaction(func(tx *gorm.DB) error {
		return b.blogDao.CreateBlog(tx, blog)
	})

}

func (b *BlogServiceImpl) PutBlog(currentUser *model.User, id int, blog *model.Blog) error {
	if currentUser == nil || !currentUser.IsAdmin {
		return fmt.Errorf("only admin can update log")
	}
	blog.ID = id
	return b.db.Transaction(func(tx *gorm.DB) error {
		return b.blogDao.UpdateBlog(tx, blog)
	})

}

func (b *BlogServiceImpl) DeleteBlog(currentUser *model.User, id int) error {
	if currentUser == nil || !currentUser.IsAdmin {
		return fmt.Errorf("only admin can delete log")
	}
	return b.db.Transaction(func(tx *gorm.DB) error {
		return b.blogDao.DeleteBlog(tx, &model.Blog{
			ID: id,
		})
	})
}
