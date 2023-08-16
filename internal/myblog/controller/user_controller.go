package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/ComradeProgrammer/blog/internal/myblog/dal/model"
	"github.com/ComradeProgrammer/blog/internal/myblog/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	BaseController
	service.UserService
}

func NewUserController(db *gorm.DB) (*UserController, error) {
	var res UserController
	baseController, err := NewBaseController(db)
	if err != nil {
		return nil, err
	}
	userService, err := service.NewUserServiceImpl(db)
	if err != nil {
		return nil, err
	}
	res.BaseController = baseController
	res.UserService = userService
	return &res, nil
}

func (u *UserController) ListUsers(c *gin.Context) {
	currentUser, err := u.BaseController.GetCurrentUserFromSession(c)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	users, err := u.UserService.ListUsers(currentUser)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, users)
}

func (u *UserController) GetUser(c *gin.Context) {
	currentUser, err := u.BaseController.GetCurrentUserFromSession(c)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("invalid parameter id: %s", c.Query("id")),
		})
		return
	}

	user, err := u.UserService.GetUser(currentUser, id)
	if err != nil {
		code := 400
		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = 404
		}
		c.JSON(code, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, user)
}

func (u *UserController) PostUser(c *gin.Context) {
	currentUser, err := u.BaseController.GetCurrentUserFromSession(c)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
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
	var user model.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = u.UserService.PostUser(currentUser, &user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "ok",
	})
}

func (u *UserController) DeleteUser(c *gin.Context) {
	currentUser, err := u.BaseController.GetCurrentUserFromSession(c)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("invalid parameter id: %s", c.Query("id")),
		})
		return

	}
	err = u.UserService.DeleteUser(currentUser, id)
	if err != nil {
		code := 400
		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = 404
		}
		c.JSON(code, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "ok",
	})
}

func (u *UserController) PutUserPassword(c *gin.Context){
	currentUser, err := u.BaseController.GetCurrentUserFromSession(c)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	username := c.Param("username")

	if username == "" {
		c.JSON(400, gin.H{
			"error":"empty username",
		})
		return
	}
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
	var data map[string]string
	err = json.Unmarshal(body, &data)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if _, ok := data["oldPassword"]; !ok {
		c.JSON(400, gin.H{
			"error": "oldPassword field must not be empty",
		})
	}
	if newPassword, ok := data["newPassword"]; !ok || newPassword == "" {
		c.JSON(400, gin.H{
			"error": "oldPassword field must not be empty",
		})
	}

	err=u.UserService.ChangePassword(currentUser,username,data["oldPassword"],data["newPassword"])


}
