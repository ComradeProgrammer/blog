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

func GetBlogs(c *gin.Context) {
	blogs, err := model.GetBlogs()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, blogs)
}

func GetBlog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("invalid parameter id: %s", c.Query("id")),
		})
		return
	}
	blog, err := model.GetBlog(id)
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
	c.JSON(200, blog)
}

func PostBlog(c *gin.Context) {
	if ok := authenticateAdmin(c); !ok {
		c.JSON(401, gin.H{
			"error": "not authorized",
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

	var blog model.Blog
	err = json.Unmarshal(body, &blog)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = model.CreateBlog(&blog)
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

func PutBlog(c *gin.Context) {
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

	body, err := ioutil.ReadAll(c.Request.Body)
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
	blog.ID = id
	err = blog.Update()
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

func DeleteBlog(c *gin.Context) {
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
	blog := model.Blog{
		ID: id,
	}
	err = blog.Delete()
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
