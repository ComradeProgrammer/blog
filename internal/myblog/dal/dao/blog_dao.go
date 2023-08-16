package dao

import (
	"gorm.io/gorm"

	"github.com/ComradeProgrammer/blog/internal/myblog/dal/model"
)

type BlogDao interface {
	GetBlogs(database *gorm.DB) ([]*model.Blog, error)
	GetBlog(database *gorm.DB, ID int) (*model.Blog, error)
	CreateBlog(database *gorm.DB, b *model.Blog) error
	UpdateBlog(database *gorm.DB, b *model.Blog) error
	DeleteBlog(database *gorm.DB, b *model.Blog) error
}

// BlogDaoImpl implements BlogDao
type BlogDaoImpl struct {
}

func (*BlogDaoImpl) GetBlogs(database *gorm.DB) ([]*model.Blog, error) {
	var res []*model.Blog
	result := database.Preload("Category").Order("create_at desc").Find(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return res, nil
}
func (*BlogDaoImpl) GetBlog(database *gorm.DB, ID int) (*model.Blog, error) {
	blog := model.Blog{
		ID: ID,
	}
	result := database.Preload("Category").Preload("Comments", func(db *gorm.DB) *gorm.DB {
		return db.Preload("User").Order("comments.create_at desc")
	}).First(&blog)

	if result.Error != nil {
		return nil, result.Error
	}
	return &blog, nil

}
func (*BlogDaoImpl) CreateBlog(database *gorm.DB, blog *model.Blog) error {
	result := database.Create(blog)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (*BlogDaoImpl) UpdateBlog(database *gorm.DB, blog *model.Blog) error {
	result := database.Where("id =  ?", blog.ID).Select("title", "content", "category_id").Updates(blog)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
func (*BlogDaoImpl) DeleteBlog(database *gorm.DB, blog *model.Blog) error {
	result := database.Delete(blog)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
