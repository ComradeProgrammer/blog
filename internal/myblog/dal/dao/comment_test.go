package dao

import (
	"log"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"

	"github.com/ComradeProgrammer/blog/internal/myblog/dal/conn"
	"github.com/ComradeProgrammer/blog/internal/myblog/dal/model"
)

func initDatabase(db *gorm.DB) {
	//insert the default admin account
	_, err := (&UserDaoImpl{}).GetUserByUserName(db, "admin")

	if err == gorm.ErrRecordNotFound {
		adminUser := &model.User{
			UserName: "admin",
			IsAdmin:  true,
		}
		adminUser.SetPassword("123456")
		err = (&UserDaoImpl{}).CreateUser(db, adminUser)
		if err != nil {
			log.Println("Unable to create admin account")
		}
	}
}

func TestCommentBasic(t *testing.T) {
	db, err := conn.ConnectSqliteDatabase("file::memory:")
	if err != nil {
		t.Errorf("failed to start test db: %v", err)
		return
	}
	var blogDao BlogDao = &BlogDaoImpl{}
	var categoryDao CategoryDao = &CategoryDaoImpl{}
	var commentDao CommentDao = &CommentDaoImpl{}
	initDatabase(db)
	Convey("TestCategoryBasic", t, func() {
		var err error
		var category = &model.Category{
			Name:        "category1",
			Description: "description1",
		}
		err = categoryDao.CreateCategory(db, category)
		So(err, ShouldBeNil)

		var blog = model.Blog{
			Title:      "Title1",
			Content:    "Content1",
			CategoryID: category.ID,
		}
		err = blogDao.CreateBlog(db, &blog)
		So(err, ShouldBeNil)

		//test create
		var comment1 = model.Comment{
			Content: "aaa",
			BlogID:  blog.ID,
			UserID:  1,
		}
		var comment2 = model.Comment{
			Content: "bbb",
			BlogID:  blog.ID,
			UserID:  1,
		}

		err = commentDao.CreateComment(db, &comment1)
		So(err, ShouldBeNil)

		err = commentDao.CreateComment(db, &comment2)
		So(err, ShouldBeNil)

		blog1, err := blogDao.GetBlog(db, blog.ID)
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

		var comment3 = model.Comment{
			Content: "bfbb",
			BlogID:  1000000,
			UserID:  1,
		}
		err = commentDao.CreateComment(db, &comment3)
		So(err, ShouldNotBeNil)

	})
}
