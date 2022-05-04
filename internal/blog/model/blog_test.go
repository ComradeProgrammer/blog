package model

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBlogBasic(t *testing.T) {
	ClearDatabase()
	Convey("TestCategoryBasic", t, func() {
		var err error
		var category = &Category{
			Name:        "category1",
			Description: "description1",
		}
		err = CreateCategory(category)
		So(err, ShouldBeNil)

		//test create
		err = CreateBlog(&Blog{
			Title:      "Title1",
			Content:    "Content1",
			CategoryID: category.ID,
		})
		So(err, ShouldBeNil)

		//test get
		blog1, err := GetBlog(0)
		So(err, ShouldBeNil)
		So(blog1, ShouldNotBeNil)
		So(blog1.Title, ShouldEqual, "Title1")
		So(blog1.Content, ShouldEqual, "Content1")
		So(blog1.Category, ShouldNotBeNil)
		So(blog1.Category.ID, ShouldEqual, category.ID)

		//test update
		blog1.Content = "modified1"
		err = blog1.Update()
		So(err, ShouldBeNil)

		blog1, err = GetBlog(0)
		So(err, ShouldBeNil)
		So(blog1.Content, ShouldEqual, "modified1")

		blogs, err := GetBlogs()
		So(err, ShouldBeNil)
		So(len(blogs), ShouldEqual, 1)

		//test delete
		err = blog1.Delete()
		So(err, ShouldBeNil)

		blogs, err = GetBlogs()
		So(err, ShouldBeNil)
		So(len(blogs), ShouldEqual, 0)

	})
}
