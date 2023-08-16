package dao

import (
	"gorm.io/gorm"

	"github.com/ComradeProgrammer/blog/internal/myblog/dal/model"
)

type CommentDao interface {
	CreateComment(database *gorm.DB, c *model.Comment) error
	Delete(database *gorm.DB, c *model.Comment) error
	GetComment(database *gorm.DB, id int) (*model.Comment, error)
}

type CommentDaoImpl struct {
}

func (*CommentDaoImpl) CreateComment(database *gorm.DB, c *model.Comment) error {
	result := database.Create(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (*CommentDaoImpl) Delete(database *gorm.DB, c *model.Comment) error {
	result := database.Delete(c)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (*CommentDaoImpl) GetComment(database *gorm.DB, id int) (*model.Comment, error) {
	comment := model.Comment{
		ID: id,
	}
	result := database.First(&comment)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}
