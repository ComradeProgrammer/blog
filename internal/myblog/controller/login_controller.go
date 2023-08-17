package controller

import (
	"encoding/json"
	"io"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginController struct {
	BaseController
}

func NewLoginController(db *gorm.DB) (*LoginController, error) {
	var res LoginController
	baseController, err := NewBaseController(db)
	if err != nil {
		return nil, err
	}
	res.BaseController = baseController
	return &res, nil
}

func (l *LoginController) Login(c *gin.Context) {
	if c.Request.Body == nil {
		c.JSON(400, gin.H{
			"error": "no body provided",
		})
		return
	}
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	var load map[string]string
	err = json.Unmarshal(body, &load)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "failed to parse Body" + err.Error(),
		})
		return
	}

	if _, ok := load["userName"]; !ok {
		c.JSON(400, gin.H{
			"error": "userName is required",
		})
		return
	}
	if _, ok := load["password"]; !ok {
		c.JSON(400, gin.H{
			"error": "password is required",
		})
		return
	}

	ok, user := l.SessionService.VerifyPassword(load["userName"], load["password"])
	if !ok {
		c.JSON(401, gin.H{
			"error": "invalid password",
		})
		return
	}

	session := sessions.Default(c)
	session.Set("userID", user.ID)
	session.Save()
	c.JSON(200, gin.H{
		"msg":  "ok",
		"user": user,
	})
}

func (l *LoginController) LogOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userID")
	c.JSON(200, gin.H{
		"msg": "ok",
	})
}
