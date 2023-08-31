package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/ComradeProgrammer/blog/internal/myblog/dal/conn"
	"github.com/ComradeProgrammer/blog/internal/myblog/dal/model"
	"github.com/ComradeProgrammer/blog/internal/myblog/service"
	"github.com/gin-gonic/gin"
)

type CommentController struct {
	BaseController
	service.CommentService
}

func NewCommentController() (*CommentController, error) {
	var res CommentController
	baseController, err := NewBaseController()
	if err != nil {
		return nil, err
	}
	commentService, err := service.NewCommentServiceImpl(conn.DB)
	if err != nil {
		return nil, err
	}
	res.BaseController = baseController
	res.CommentService = commentService
	return &res, nil
}

func (cc *CommentController) PostComment(c *gin.Context) {
	currentUser, err := cc.BaseController.GetCurrentUserFromSession(c)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
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

	var comment model.Comment
	err = json.Unmarshal(body, &comment)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = cc.CommentService.PostComment(currentUser, &comment)

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

func (cc *CommentController) DeleteComment(c *gin.Context) {
	currentUser, err := cc.BaseController.GetCurrentUserFromSession(c)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	if c.Request.Body == nil {
		c.JSON(400, gin.H{
			"error": "no body provided",
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
	err = cc.CommentService.DeleteComment(currentUser, id)

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
