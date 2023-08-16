package service

import (
	"github.com/ComradeProgrammer/blog/internal/myblog/dal/dao"
)

type ServiceBase struct {
	blogDao     dao.BlogDao
	categoryDao dao.CategoryDao
	commentDao  dao.CommentDao
	userDao     dao.UserDao
}

func NewServiceBase() ServiceBase {
	return ServiceBase{
		blogDao:     &dao.BlogDaoImpl{},
		categoryDao: &dao.CategoryDaoImpl{},
		commentDao:  &dao.CommentDaoImpl{},
		userDao:     &dao.UserDaoImpl{},
	}
}
