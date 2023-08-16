package dao

import (
	"gorm.io/gorm"

	"github.com/ComradeProgrammer/blog/internal/myblog/dal/model"
)

type CategoryDao interface {
	GetCategories(database *gorm.DB) ([]*model.Category, error)
	GetCategory(database *gorm.DB, ID int) (*model.Category, error)
	CreateCategory(database *gorm.DB, c *model.Category) error
	Update(database *gorm.DB, c *model.Category) error
	Delete(database *gorm.DB, c *model.Category) error
}
type CategoryDaoImpl struct {
}

func (*CategoryDaoImpl) GetCategories(database *gorm.DB) ([]*model.Category, error) {
	//No preload blogs
	var res []*model.Category
	result := database.Order("create_at desc").Find(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return res, nil
}
func (*CategoryDaoImpl) GetCategory(database *gorm.DB, ID int) (*model.Category, error) {
	// blogs will be preloaded
	category := model.Category{
		ID: ID,
	}
	result := database.Preload("Blogs", func(db *gorm.DB) *gorm.DB {
		return db.Order("blogs.create_at desc")
	}).First(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil

}
func (*CategoryDaoImpl) CreateCategory(database *gorm.DB, c *model.Category) error {
	result := database.Create(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (*CategoryDaoImpl) Update(database *gorm.DB, c *model.Category) error {
	result := database.Where("id =  ?", c.ID).Select("name", "description").Updates(c)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
func (*CategoryDaoImpl) Delete(database *gorm.DB, c *model.Category) error {
	result := database.Delete(c)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
