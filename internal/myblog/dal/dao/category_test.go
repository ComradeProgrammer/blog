package dao

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/ComradeProgrammer/blog/internal/myblog/dal/conn"
	"github.com/ComradeProgrammer/blog/internal/myblog/dal/model"
)

func TestCategoryBasic(t *testing.T) {
	Convey("TestCategoryBasic", t, func() {
		db, err := conn.ConnectSqliteDatabase("file::memory:")
		So(err, ShouldBeNil)
		var categoryDao CategoryDao = &CategoryDaoImpl{}
		//test create
		category := &model.Category{
			Name:        "category1",
			Description: "description1",
		}
		err = categoryDao.CreateCategory(db, category)
		So(err, ShouldBeNil)

		//test get
		category1, err := categoryDao.GetCategory(db, category.ID)
		So(err, ShouldBeNil)
		So(category1, ShouldNotBeNil)
		So(category1.Name, ShouldEqual, "category1")
		So(category1.Description, ShouldEqual, "description1")

		//test update
		categoryForUpdate := &model.Category{
			ID:          category1.ID,
			Name:        "modified1",
			Description: "description1",
		}
		err = categoryDao.Update(db, categoryForUpdate)
		So(err, ShouldBeNil)

		category2, err := categoryDao.GetCategory(db, category.ID)
		So(err, ShouldBeNil)
		So(category2.Name, ShouldEqual, "modified1")
		So(category2.Description, ShouldEqual, categoryForUpdate.Description)
		So(category2.CreateAt.Equal(category1.CreateAt), ShouldBeTrue)

		//test get all
		categories, err := categoryDao.GetCategories(db)
		So(err, ShouldBeNil)
		So(len(categories), ShouldEqual, 1)

		//test delete
		err = categoryDao.Delete(db, category1)
		So(err, ShouldBeNil)

		categories, err = categoryDao.GetCategories(db)
		So(err, ShouldBeNil)
		So(len(categories), ShouldEqual, 0)

	})
}

func TestCategoryAdvanced(t *testing.T) {
	var blog1 *model.Blog
	Convey("TestCategoryAdvanced", t, func() {

		db, err := conn.ConnectSqliteDatabase("file::memory:")
		So(err, ShouldBeNil)
		var blogDao BlogDao = &BlogDaoImpl{}
		var categoryDao CategoryDao = &CategoryDaoImpl{}

		var category *model.Category
		Convey("TestCategoryPreload", func() {
			var err error
			category = &model.Category{
				Name:        "category1",
				Description: "description1",
			}
			err = categoryDao.CreateCategory(db, category)
			So(err, ShouldBeNil)

			blog1 = &model.Blog{
				Title:      "Title1",
				Content:    "Content1",
				CategoryID: category.ID,
			}
			err = blogDao.CreateBlog(db, blog1)
			So(err, ShouldBeNil)

			//GetCategories() should not have preloads
			categories, err := categoryDao.GetCategories(db)
			So(err, ShouldBeNil)
			So(len(categories), ShouldEqual, 1)
			So(categories[0].Blogs, ShouldBeNil)

			//GetCategory() should have preloads
			category, err = categoryDao.GetCategory(db, category.ID)
			So(err, ShouldBeNil)
			So(category.Blogs, ShouldNotBeNil)
			So(len(*category.Blogs), ShouldEqual, 1)
			So((*category.Blogs)[0].ID, ShouldEqual, blog1.ID)
		})

		Convey("TestForeignKeyConstraint", func() {
			err = categoryDao.Delete(db, category)

			So(err, ShouldNotBeNil)

			blog1.CategoryID += 10000
			err = blogDao.UpdateBlog(db, blog1)
			So(err, ShouldNotBeNil)
		})
	})

}
