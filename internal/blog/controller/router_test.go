package controller

import (
	"io/ioutil"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/ComradeProgrammer/blog/internal/blog/model"
)

func restoreTestDatabase() {
	model.ClearDatabase()
	model.InitDatabase()
	//Generate test data
	category1 := &model.Category{
		Name:        "category1",
		Description: "description1",
	}
	model.CreateCategory(category1)

	category2 := &model.Category{
		Name:        "category2",
		Description: "description2",
	}
	model.CreateCategory(category2)

	blog1 := &model.Blog{
		Title:      "title1",
		Content:    "content1",
		CategoryID: category1.ID,
	}
	model.CreateBlog(blog1)
	blog2 := &model.Blog{
		Title:      "title2",
		Content:    "content2",
		CategoryID: category1.ID,
	}
	model.CreateBlog(blog2)

	blog3 := &model.Blog{
		Title:      "title3",
		Content:    "content3",
		CategoryID: category2.ID,
	}
	model.CreateBlog(blog3)
	blog4 := &model.Blog{
		Title:      "title4",
		Content:    "content4",
		CategoryID: category2.ID,
	}
	model.CreateBlog(blog4)
}

func TestMain(m *testing.M) {
	model.ConnectDatabase("../../../database_test.sqlite")
	model.InitDatabase()
	restoreTestDatabase();
	//shut down output
	gin.DefaultWriter = ioutil.Discard
	m.Run()
}
