package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          int       `gorm:"id;primaryKey;autoIncrement"`
	Name        string    `gorm:"name" json:"name"`
	Description string    `gorm:"description" json:"description"`
	CreateAt    time.Time `gorm:"create_at;autoCreateTime" json:"createAt"`
	UpdateAt    time.Time `gorm:"update_at;autoUpdateTime" json:"updateAt"`

	Blogs *[]Blog `gorm:"foreignKey:category_id" json:"blogs"`
}

//No preload blogs
func GetCategories() ([]*Category, error) {
	var res []*Category
	result := database.Order("create_at desc").Find(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return res, nil

}

//blogs will be preloaded
func GetCategory(ID int) (*Category, error) {
	category := Category{
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

func CreateCategory(c *Category) error {
	result := database.Create(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *Category) Update() error {
	result := database.Where("id =  ?", c.ID).Select("name", "description").Updates(c)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (c *Category) Delete() error {
	result := database.Delete(c)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
