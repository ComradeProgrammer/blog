package model

import (
	"time"
)

type Category struct {
	ID          int       `gorm:"id;primaryKey;autoIncrement"`
	Name        string    `gorm:"name" json:"name"`
	Description string    `gorm:"description" json:"description"`
	CreateAt    time.Time `gorm:"create_at;autoCreateTime" json:"createAt"`
	UpdateAt    time.Time `gorm:"update_at;autoUpdateTime" json:"updateAt"`

	Blogs *[]Blog `gorm:"foreignKey:category_id" json:"blogs"`
}
