package controller

import (
	"encoding/json"
	"io/ioutil"

	"github.com/ComradeProgrammer/blog/internal/blog/model"
	"github.com/gin-gonic/gin"
)

func PostComment(c *gin.Context) {
	user := getUserFromSession(c)
	if user == nil {
		c.JSON(401, gin.H{
			"error": "You need to login",
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

	var comment model.Comment
	err = json.Unmarshal(body, &comment)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if comment.Content == "" {
		c.JSON(400, gin.H{
			"error": "comment must not be empty",
		})
		return
	}

	if comment.UserID != user.ID {
		c.JSON(400, gin.H{
			"error": "comment.UserID is not current user",
		})
	}

	err = model.CreateComment(&comment)
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
