package controller

import (
	"github.com/ComradeProgrammer/blog/internal/blog/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func authenticateAdmin(c *gin.Context) bool {
	session := sessions.Default(c)
	if session.Get("userID") == nil {
		return false
	}
	id := session.Get("userID").(int)
	user, err := model.GetUserByID(id)
	if err != nil {
		return false
	}
	return user.IsAdmin
}

func getUserFromSession(c *gin.Context) *model.User{
	session := sessions.Default(c)
	if session.Get("userID") == nil {
		return nil
	}
	id := session.Get("userID").(int)
	user, err := model.GetUserByID(id)
	if err != nil {
		return nil
	}
	return user
}
