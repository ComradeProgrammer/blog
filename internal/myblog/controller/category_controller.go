package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"gorm.io/gorm"

	"github.com/ComradeProgrammer/blog/internal/myblog/dal/model"
	"github.com/ComradeProgrammer/blog/internal/myblog/service"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	BaseController
	service.CategoryService
}

func NewCategoryController(db *gorm.DB) (*CategoryController, error) {
	var res CategoryController
	baseController, err := NewBaseController(db)
	if err != nil {
		return nil, err
	}
	CategoryService, err := service.NewCategoryServiceImpl(db)
	if err != nil {
		return nil, err
	}
	res.BaseController = baseController
	res.CategoryService = CategoryService
	return &res, nil
}

func (cc *CategoryController) ListCategory(c *gin.Context) {
	currentUser, err := cc.BaseController.GetCurrentUserFromSession(c)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	categories, err := cc.ListCategories(currentUser)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, categories)
}

func (cc *CategoryController) GetCategory(c *gin.Context) {
	currentUser, err := cc.BaseController.GetCurrentUserFromSession(c)
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
	category, err := cc.CategoryService.GetCategory(currentUser, id)
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
	c.JSON(200, category)
}

func (cc *CategoryController) PostCategory(c *gin.Context) {
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
	var category model.Category
	err = json.Unmarshal(body, &category)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = cc.CategoryService.PostCategory(currentUser, &category)
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

func (cc *CategoryController) PutCategory(c *gin.Context) {
	currentUser, err := cc.BaseController.GetCurrentUserFromSession(c)
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
	var category model.Category
	err = json.Unmarshal(body, &category)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = cc.CategoryService.PutCategory(currentUser, id, &category)
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

func (cc *CategoryController) DeleteCategory(c *gin.Context) {
	currentUser, err := cc.BaseController.GetCurrentUserFromSession(c)
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
	err = cc.CategoryService.DeleteCategory(currentUser, id)
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
