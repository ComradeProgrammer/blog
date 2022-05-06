package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/ComradeProgrammer/blog/internal/blog/model"
)

func GetCategories(c *gin.Context) {
	categories, err := model.GetCategories()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, categories)
}

func GetCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("invalid parameter id: %s", c.Query("id")),
		})
		return
	}
	category, err := model.GetCategory(id)
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

func PostCategory(c *gin.Context) {
	if ok := authenticateAdmin(c); !ok {
		c.JSON(401, gin.H{
			"error": "not authorized",
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

	var category model.Category
	err = json.Unmarshal(body, &category)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if category.Name == "" {
		c.JSON(400, gin.H{
			"error": "name field must not be empty",
		})
	}

	err = model.CreateCategory(&category)
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

func PutCategory(c *gin.Context) {
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

	var category model.Category
	err = json.Unmarshal(body, &category)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if category.Name == "" {
		c.JSON(400, gin.H{
			"error": "name field must not be empty",
		})
	}
	category.ID = id

	err = category.Update()
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

func DeleteCategory(c *gin.Context) {
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
	category := model.Category{
		ID: id,
	}
	err = category.Delete()
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
