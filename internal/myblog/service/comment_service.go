package service

import (
	"fmt"

	"github.com/ComradeProgrammer/blog/internal/myblog/dal/model"
	"gorm.io/gorm"
)

type CommentService interface {
	PostComment(currentUser *model.User, comment *model.Comment) error
	DeleteComment(currentUser *model.User, id int) error
}

type CommentServiceImpl struct {
	ServiceBase
	db *gorm.DB
}

func NewCommentServiceImpl(db *gorm.DB) (*CommentServiceImpl, error) {
	return &CommentServiceImpl{
		ServiceBase: NewServiceBase(),
		db:          db,
	}, nil
}

func (c *CommentServiceImpl) PostComment(currentUser *model.User, comment *model.Comment) error {
	if currentUser == nil {
		return fmt.Errorf("login is required")
	}
	if comment.UserID != currentUser.ID {
		return fmt.Errorf("comment.UserID is not current user")
	}
	if comment.Content == "" {
		return fmt.Errorf("comment must not be empty")
	}

	return c.db.Transaction(func(tx *gorm.DB) error {
		return c.commentDao.CreateComment(tx, comment)
	})
}
func (c *CommentServiceImpl) DeleteComment(currentUser *model.User, id int) error {
	if currentUser == nil {
		return fmt.Errorf("login is required")
	}

	return c.db.Transaction(func(tx *gorm.DB) error {
		comment, err := c.commentDao.GetComment(tx, id)
		if err != nil {
			return err
		}
		if comment.UserID != currentUser.ID {
			return fmt.Errorf("comment.UserID is not current user")
		}

		return c.commentDao.Delete(tx, &model.Comment{
			ID: id,
		})
	})
}
