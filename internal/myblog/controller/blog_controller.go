package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/ComradeProgrammer/blog/internal/myblog/dal/conn"
	"github.com/ComradeProgrammer/blog/internal/myblog/dal/model"
	"github.com/ComradeProgrammer/blog/internal/myblog/service"
)

type BlogController struct {
	BaseController
	service.BlogService
}

func NewBlogController() (*BlogController, error) {
	var res BlogController
	baseController, err := NewBaseController()
	if err != nil {
		return nil, err
	}
	blogService, err := service.NewBlogServiceImpl(conn.DB)
	if err != nil {
		return nil, err
	}
	res.BaseController = baseController
	res.BlogService = blogService
	return &res, nil
}

func (b *BlogController) ListBlogs(c *gin.Context) {
	currentUser, err := b.BaseController.GetCurrentUserFromSession(c)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	blogs, err := b.BlogService.ListBlogs(currentUser)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, blogs)
}

func (b *BlogController) GetBlog(c *gin.Context) {
	currentUser, err := b.BaseController.GetCurrentUserFromSession(c)
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
	blog, err := b.BlogService.GetBlog(currentUser, id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, blog)
}

func (b *BlogController) PostBlog(c *gin.Context) {
	currentUser, err := b.BaseController.GetCurrentUserFromSession(c)
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
	var blog model.Blog
	err = json.Unmarshal(body, &blog)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = b.BlogService.PostBlog(currentUser, &blog)
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

func (b *BlogController) PutBlog(c *gin.Context) {
	currentUser, err := b.BaseController.GetCurrentUserFromSession(c)
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
	var blog model.Blog
	err = json.Unmarshal(body, &blog)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = b.BlogService.PutBlog(currentUser, id, &blog)
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

func (b *BlogController) DeleteBlog(c *gin.Context) {
	currentUser, err := b.BaseController.GetCurrentUserFromSession(c)
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
	err = b.BlogService.DeleteBlog(currentUser, id)
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
