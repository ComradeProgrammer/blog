package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/ComradeProgrammer/blog/internal/blog/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUsers(c *gin.Context) {
	users, err := model.GetUsers()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, users)
}

func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("invalid parameter id: %s", c.Query("id")),
		})
		return
	}
	user, err := model.GetUserByID(id)
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

func PostUser(c *gin.Context) {
	if c.Request.Body == nil {
		c.JSON(400, gin.H{
			"error": "no body provided",
		})
		return
	}
	body, err := ioutil.ReadAll(c.Request.Body)
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

	_, err = model.GetUserByUserName(user.UserName)
	if err != gorm.ErrRecordNotFound {
		c.JSON(400, gin.H{
			"error": "User Has been Created",
		})
		return
	}

	password := user.Password
	user.SetPassword(password)

	err = model.CreateUser(&user)
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

func DeleteUser(c *gin.Context) {
	if ok := authenticateAdmin(c); !ok {
		c.JSON(401, gin.H{
			"error": "not authorized",
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
	if id==1{
		c.JSON(400, gin.H{
			"error": "you cannot delete initial user",
		})
		return
	}
	user := model.User{
		ID: id,
	}
	err = user.Delete()
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

func PutUserPassword(c *gin.Context) {
	username := c.Param("username")

	if username == "" {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("empty username"),
		})
		return
	}
	if c.Request.Body == nil {
		c.JSON(400, gin.H{
			"error": "no body provided",
		})
		return
	}
	body, err := ioutil.ReadAll(c.Request.Body)
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
	user, err := model.GetUserByUserName(username)
	if err != nil {
		msg := err.Error()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg = "user not found"
		}
		c.JSON(400, gin.H{
			"error": msg,
		})
		return
	}

	if ok := user.VerifyPassword(data["oldPassword"]); !ok {
		c.JSON(400, gin.H{
			"error": "incorrect password",
		})
		return
	}

	user.SetPassword(data["newPassword"])
	err = user.Update()
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
