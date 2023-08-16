package model

import (
	"time"
)

type Comment struct {
	ID       int       `gorm:"id;primaryKey;autoIncrement"`
	Content  string    `gorm:"content" json:"content"`
	CreateAt time.Time `gorm:"create_at;autoCreateTime" json:"createAt"`
	UpdateAt time.Time `gorm:"update_at;autoUpdateTime" json:"updateAt"`

	//Blog   *Blog `gorm:"foreignKey:blog_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"blog"`
	BlogID int `gorm:"blog_id" json:"blogID"`

	User   *User `gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	UserID int   `gorm:"user_id" json:"userID"`
}
