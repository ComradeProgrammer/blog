package model

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCategoryBasic(t *testing.T) {
	ClearDatabase()
	Convey("TestCategoryBasic", t, func() {
		var err error
		//test create
		category := &Category{
			Name:        "category1",
			Description: "description1",
		}
		err = CreateCategory(category)
		So(err, ShouldBeNil)

		//test get
		category1, err := GetCategory(category.ID)
		So(err, ShouldBeNil)
		So(category1, ShouldNotBeNil)
		So(category1.Name, ShouldEqual, "category1")
		So(category1.Description, ShouldEqual, "description1")

		//test update
		categoryForUpdate := &Category{
			ID: category1.ID,
			Name:        "modified1",
			Description: "description1",
		}
		err = categoryForUpdate.Update()
		So(err, ShouldBeNil)

		category2, err := GetCategory(category.ID)
		So(err, ShouldBeNil)
		So(category2.Name, ShouldEqual, "modified1")
		So(category2.Description,ShouldEqual,categoryForUpdate.Description)
		So(category2.CreateAt.Equal(category1.CreateAt),ShouldBeTrue)

		//test get all
		categories, err := GetCategories()
		So(err, ShouldBeNil)
		So(len(categories), ShouldEqual, 1)

		//test delete
		err = category1.Delete()
		So(err, ShouldBeNil)

		categories, err = GetCategories()
		So(err, ShouldBeNil)
		So(len(categories), ShouldEqual, 0)

	})
}

func TestCategoryAdvanced(t *testing.T) {
	ClearDatabase()
	var category *Category
	var blog1 *Blog
	Convey("TestCategoryPreload", t, func() {
		var err error
		category = &Category{
			Name:        "category1",
			Description: "description1",
		}
		err = CreateCategory(category)
		So(err, ShouldBeNil)

		blog1 = &Blog{
			Title:      "Title1",
			Content:    "Content1",
			CategoryID: category.ID,
		}
		err = CreateBlog(blog1)

		//GetCategories() should not have preloads
		categories, err := GetCategories()
		So(err, ShouldBeNil)
		So(len(categories), ShouldEqual, 1)
		So(categories[0].Blogs, ShouldBeNil)

		//GetCategory() should have preloads
		category, err = GetCategory(category.ID)
		So(err, ShouldBeNil)
		So(category.Blogs, ShouldNotBeNil)
		So(len(*category.Blogs), ShouldEqual, 1)
		So((*category.Blogs)[0].ID, ShouldEqual, blog1.ID)
	})

	Convey("TestForeignKeyConstraint", t, func() {
		err := category.Delete()
		So(err, ShouldNotBeNil)

		blog1.CategoryID += 10000
		err = blog1.Update()
		So(err, ShouldNotBeNil)
	})

}
