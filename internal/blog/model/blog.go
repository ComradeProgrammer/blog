package model

import (
	"time"
)

type Blog struct {
	ID       int       `gorm:"id;primaryKey;autoIncrement"`
	Title    string    `gorm:"title"`
	Content  string    `gorm:"content"`
	CreateAt time.Time `gorm:"create_at;autoCreateTime"`
	UpdateAt time.Time `gorm:"update_at;autoUpdateTime"`

	Category   *Category `gorm:"foreignKey:category_id;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	CategoryID int       `gorm:"category_id"`
}

func GetBlogs() ([]*Blog, error) {
	var res []*Blog
	result := database.Preload("Category").Find(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return res, nil
}

func GetBlog(ID int) (*Blog, error) {
	blog := Blog{
		ID: ID,
	}
	result := database.Preload("Category").Find(&blog)
	if result.Error != nil {
		return nil, result.Error
	}
	return &blog, nil
}

func CreateBlog(b *Blog) error {
	result := database.Create(b)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (b *Blog) Update() error {
	result := database.Save(b)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (b *Blog) Delete() error {
	result := database.Delete(b)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
