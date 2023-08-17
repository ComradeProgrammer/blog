package model

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCommentBasic(t *testing.T) {
	ClearDatabase()
	InitDatabase()
	Convey("TestCategoryBasic", t, func() {
		var err error
		var category = &Category{
			Name:        "category1",
			Description: "description1",
		}
		err = CreateCategory(category)
		So(err, ShouldBeNil)

		var blog = Blog{
			Title:      "Title1",
			Content:    "Content1",
			CategoryID: category.ID,
		}
		err = CreateBlog(&blog)
		So(err, ShouldBeNil)

		//test create
		var comment1 = Comment{
			Content: "aaa",
			BlogID:  blog.ID,
			UserID:  1,
		}
		var comment2 = Comment{
			Content: "bbb",
			BlogID:  blog.ID,
			UserID:  1,
		}

		err = CreateComment(&comment1)
		So(err, ShouldBeNil)

		err = CreateComment(&comment2)
		So(err, ShouldBeNil)

		blog1, err := GetBlog(blog.ID)
		So(err, ShouldBeNil)
		So(blog1, ShouldNotBeNil)
		So(blog1.Title, ShouldEqual, "Title1")
		So(blog1.Content, ShouldEqual, "Content1")
		So(blog1.Category, ShouldNotBeNil)
		So(blog1.Category.ID, ShouldEqual, category.ID)
		So(len(*blog1.Comments), ShouldEqual, 2)

		So((*blog1.Comments)[0].ID, ShouldEqual, comment2.ID)
		So((*blog1.Comments)[1].ID, ShouldEqual, comment1.ID)

		So((*blog1.Comments)[0].UserID, ShouldEqual, 1)
		So((*blog1.Comments)[0].User.UserName, ShouldEqual, "admin")

		var comment3 = Comment{
			Content: "bfbb",
			BlogID:  1000000,
			UserID:  1,
		}
		err = CreateComment(&comment3)
		So(err, ShouldNotBeNil)

	})
}
