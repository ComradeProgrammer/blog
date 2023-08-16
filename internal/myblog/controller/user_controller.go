package controller

import (
	"github.com/ComradeProgrammer/blog/internal/myblog/service"
	"gorm.io/gorm"
)

type UserController struct {
	BaseController
	service.UserService
}


func NewUserController(db *gorm.DB) (UserController, error){
	var res UserController
	baseController, err := NewBaseController(db)
	if err != nil {
		return res, err
	}
	userService, err := service.NewUserServiceImpl(db)
	if err != nil {
		return res, err
	}
	res.BaseController = baseController
	res.UserService = userService
	return res, nil
}

