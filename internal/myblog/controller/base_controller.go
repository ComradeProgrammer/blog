package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/ComradeProgrammer/blog/internal/myblog/dal/conn"
	"github.com/ComradeProgrammer/blog/internal/myblog/dal/model"
	"github.com/ComradeProgrammer/blog/internal/myblog/service"
)

type BaseController struct {
	db *gorm.DB
	service.SessionService
}

func NewBaseController() (BaseController, error) {
	sessionSvc, err := service.NewSessionService(conn.DB)
	if err != nil {
		return BaseController{}, err
	}
	return BaseController{
		db:             conn.DB,
		SessionService: sessionSvc,
	}, nil
}

func (b *BaseController) GetCurrentUserFromSession(c *gin.Context) (*model.User, error) {
	session := sessions.Default(c)
	if session.Get("userID") == nil {
		return nil, nil
	}
	id := session.Get("userID").(int)

	return b.SessionService.GetUser(id)
}
