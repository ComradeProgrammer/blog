package dao

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/ComradeProgrammer/blog/internal/myblog/dal/conn"
	"github.com/ComradeProgrammer/blog/internal/myblog/dal/model"
)

func TestBlogBasic(t *testing.T) {
	Convey("TestCategoryBasic", t, func() {
		db, err := conn.ConnectSqliteDatabase("file::memory:")
		So(err, ShouldBeNil)
		var blogDao BlogDao = &BlogDaoImpl{}
		var categoryDao CategoryDao = &CategoryDaoImpl{}

		var category = &model.Category{
			Name:        "category1",
			Description: "description1",
		}
		err = categoryDao.CreateCategory(db, category)
		So(err, ShouldBeNil)

		//test create
		err = blogDao.CreateBlog(db, &model.Blog{
			Title:      "Title1",
			Content:    "Content1",
			CategoryID: category.ID,
		})
		So(err, ShouldBeNil)

		//test get
		blog1, err := blogDao.GetBlog(db, 0)
		So(err, ShouldBeNil)
		So(blog1, ShouldNotBeNil)
		So(blog1.Title, ShouldEqual, "Title1")
		So(blog1.Content, ShouldEqual, "Content1")
		So(blog1.Category, ShouldNotBeNil)
		So(blog1.Category.ID, ShouldEqual, category.ID)

		//test update
		blog1.Content = "modified1"
		err = blogDao.UpdateBlog(db, blog1)
		So(err, ShouldBeNil)

		blog1, err = blogDao.GetBlog(db, 0)
		So(err, ShouldBeNil)
		So(blog1.Content, ShouldEqual, "modified1")

		blogs, err := blogDao.GetBlogs(db)
		So(err, ShouldBeNil)
		So(len(blogs), ShouldEqual, 1)

		//test delete
		err = blogDao.DeleteBlog(db, blog1)
		So(err, ShouldBeNil)

		blogs, err = blogDao.GetBlogs(db)
		So(err, ShouldBeNil)
		So(len(blogs), ShouldEqual, 0)

	})
}
