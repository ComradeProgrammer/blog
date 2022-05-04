package model

import (
	"time"
)

type Category struct {
	ID          int       `gorm:"id;primaryKey;autoIncrement"`
	Name        string    `gorm:"name"`
	Description string    `gorm:"description"`
	CreateAt    time.Time `gorm:"create_at;autoCreateTime"`
	UpdateAt    time.Time `gorm:"update_at;autoUpdateTime"`

	Blogs *[]Blog `gorm:"foreignKey:category_id"`
}

//No preload blogs
func GetCategories() ([]*Category, error) {
	var res []*Category
	result := database.Find(&res)
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
	result := database.Preload("Blogs").Find(&category)
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
	result := database.Save(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *Category) Delete() error {
	result := database.Delete(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
