package model

import (
	"time"
)

type Blog struct {
	ID       int       `gorm:"id;primaryKey;autoIncrement"`
	Title    string    `gorm:"title" json:"title"`
	Content  string    `gorm:"content" json:"content"`
	CreateAt time.Time `gorm:"create_at;autoCreateTime" json:"createAt"`
	UpdateAt time.Time `gorm:"update_at;autoUpdateTime" json:"updateAt"`

	Category   *Category `gorm:"foreignKey:category_id;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;" json:"category"`
	CategoryID int       `gorm:"category_id" json:"categoryID"`

	Comments *[]Comment `gorm:"foreignKey:blog_id" json:"comments"`
}
