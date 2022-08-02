package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID       int       `gorm:"id;primaryKey;autoIncrement"`
	Content  string    `gorm:"content" json:"content"`
	CreateAt time.Time `gorm:"create_at;autoCreateTime" json:"createAt"`
	UpdateAt time.Time `gorm:"update_at;autoUpdateTime" json:"updateAt"`

	//Blog   *Blog `gorm:"foreignKey:blog_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"blog"`
	BlogID int   `gorm:"blog_id" json:"blogID"`

	User   *User `gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	UserID int   `gorm:"user_id" json:"userID"`
}

func CreateComment(c *Comment) error {
	result := database.Create(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *Comment) Delete() error {
	result := database.Delete(c)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
