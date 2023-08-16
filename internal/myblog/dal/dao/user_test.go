package dao

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/ComradeProgrammer/blog/internal/myblog/dal/conn"
	"github.com/ComradeProgrammer/blog/internal/myblog/dal/model"
)

func TestPassword(t *testing.T) {
	Convey("TestPassword", t, func() {
		u := model.User{}
		err := u.SetPassword("password")
		So(err, ShouldBeNil)

		ok := u.VerifyPassword("password")
		So(ok, ShouldBeTrue)
		ok = u.VerifyPassword("passwordd")
		So(ok, ShouldBeFalse)

		err = u.SetPassword("")
		So(err, ShouldNotBeNil)
	})
}

func TestUserBasic(t *testing.T) {
	db, err := conn.ConnectSqliteDatabase("file::memory:")
	if err != nil {
		t.Errorf("failed to start test db: %v", err)
		return
	}
	var userDao UserDao = &UserDaoImpl{}
	Convey("TestUserBasic", t, func() {
		var err error
		var user = &model.User{
			UserName: "abc@qq.com",
			IsAdmin:  false,
		}
		user.SetPassword("password")
		//test create
		err = userDao.CreateUser(db, user)
		So(err, ShouldBeNil)

		//test get
		userCopy2, err := userDao.GetUserByUserName(db, "abc@qq.com")
		So(err, ShouldBeNil)
		So(userCopy2, ShouldNotBeNil)
		So(user.UserName, ShouldEqual, userCopy2.UserName)
		So(user.IsAdmin, ShouldEqual, userCopy2.IsAdmin)

		userCopy1, err := userDao.GetUserByID(db, user.ID)
		So(err, ShouldBeNil)
		So(userCopy1, ShouldNotBeNil)
		So(user.UserName, ShouldEqual, userCopy1.UserName)
		So(user.IsAdmin, ShouldEqual, userCopy1.IsAdmin)
		So(userCopy1.VerifyPassword("password"), ShouldBeTrue)
		So(userCopy1.VerifyPassword("passwdord"), ShouldBeFalse)

		users, err := userDao.GetUsers(db)
		So(err, ShouldBeNil)
		So(len(users), ShouldEqual, 1)

		//test update

		userForUpdate := &model.User{
			ID:       user.ID,
			UserName: "abcd@qq.com",
			IsAdmin:  user.IsAdmin,
		}
		err = userDao.Update(db, userForUpdate)
		So(err, ShouldBeNil)

		userCopy3, err := userDao.GetUserByID(db, user.ID)
		So(err, ShouldBeNil)
		So(userCopy3.UserName, ShouldEqual, userForUpdate.UserName)
		//So(userCopy3.CreateAt, ShouldEqual, user.CreateAt)

		//test dupliacte insert
		var user4 = &model.User{
			UserName: "abcd@qq.com",
			IsAdmin:  false,
		}
		user.SetPassword("password2")
		//test create
		err = userDao.CreateUser(db, user4)
		So(err, ShouldNotBeNil)

		//test delete
		err = userDao.Delete(db, user)

		So(err, ShouldBeNil)
		users, err = userDao.GetUsers(db)
		So(err, ShouldBeNil)
		So(len(users), ShouldEqual, 0)

	})
}
