package model

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPassword(t *testing.T) {
	Convey("TestPassword", t, func() {
		u := User{}
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
	ClearDatabase()
	Convey("TestUserBasic", t, func() {
		var err error
		var user = &User{
			UserName: "abc@qq.com",
			IsAdmin:  false,
		}
		user.SetPassword("password")
		//test create
		err = CreateUser(user)
		So(err, ShouldBeNil)

		//test get
		userCopy1, err := GetUserByID(user.ID)
		So(err, ShouldBeNil)
		So(userCopy1, ShouldNotBeNil)
		So(reflect.DeepEqual(*user, *userCopy1), ShouldBeTrue)
		So(userCopy1.VerifyPassword("password"), ShouldBeTrue)
		So(userCopy1.VerifyPassword("passwdord"), ShouldBeFalse)

		userCopy2, err := GetUserByUserName(user.UserName)
		So(err, ShouldBeNil)
		So(userCopy2, ShouldNotBeNil)
		So(reflect.DeepEqual(*user, *userCopy2), ShouldBeTrue)

		users, err := GetUsers()
		So(err, ShouldBeNil)
		So(len(users), ShouldEqual, 1)

		//test update
		user.UserName = "abcd@qq.com"
		err = user.Update()
		So(err, ShouldBeNil)

		userCopy3, err := GetUserByID(user.ID)
		So(err, ShouldBeNil)
		So(userCopy3.UserName, ShouldEqual, user.UserName)

		//test dupliacte insert
		var user4 = &User{
			UserName: "abcd@qq.com",
			IsAdmin:  false,
		}
		user.SetPassword("password2")
		//test create
		err = CreateUser(user4)
		So(err, ShouldNotBeNil)

		//test delete
		err = user.Delete()
		So(err, ShouldBeNil)
		users, err = GetUsers()
		So(err, ShouldBeNil)
		So(len(users), ShouldEqual, 0)

	})
}
