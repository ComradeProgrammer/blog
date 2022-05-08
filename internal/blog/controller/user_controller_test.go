package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ComradeProgrammer/blog/internal/blog/model"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetUsers(t *testing.T) {
	restoreTestDatabase()
	r := GetGinEngine()
	Convey("TestGetUsers", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/user", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 200)
		body, err := io.ReadAll(w.Body)
		So(err, ShouldBeNil)
		var res []model.User
		err = json.Unmarshal(body, &res)
		So(err, ShouldBeNil)

		So(len(res), ShouldEqual, 1)
		So(res[0].UserName, ShouldEqual, "admin")
		So(res[0].IsAdmin, ShouldEqual, true)
		So(res[0].VerifyPassword("123456"), ShouldEqual, true)
	})
}

func TestGetUser(t *testing.T) {
	restoreTestDatabase()
	r := GetGinEngine()
	Convey("TestGetCategory", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/user/1", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 200)
		body, err := io.ReadAll(w.Body)
		So(err, ShouldBeNil)

		var res model.User
		err = json.Unmarshal(body, &res)
		So(err, ShouldBeNil)

		So(res.UserName, ShouldEqual, "admin")
		So(res.IsAdmin, ShouldEqual, true)
		So(res.VerifyPassword("123456"), ShouldEqual, true)

	})
	Convey("TestGetUser404", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/user/100", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 404)
	})
}

func TestPostUser(t *testing.T) {
	restoreTestDatabase()
	r := GetGinEngine()
	Convey("TestPostUserEmptyBody", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/user", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 400)
	})

	Convey("TestPostUser", t, func() {
		w := httptest.NewRecorder()
		cookie := getAdminCookie(r)

		user := model.User{
			UserName: "category3",
			IsAdmin:  false,
		}
		user.SetPassword("123456789")
		data, _ := json.Marshal(user)

		req, _ := http.NewRequest("POST", "/api/user", bytes.NewBuffer(data))
		req.Header.Set("Cookie", cookie)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 200)

		w = httptest.NewRecorder()
		req, _ = http.NewRequest("GET", "/api/user/2", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 200)
		body, err := io.ReadAll(w.Body)
		So(err, ShouldBeNil)

		var res model.User
		err = json.Unmarshal(body, &res)
		So(err, ShouldBeNil)

		So(res.UserName, ShouldEqual, "category3")
		So(res.IsAdmin, ShouldEqual, false)
		So(res.VerifyPassword("123456789"), ShouldEqual, true)

	})

	Convey("TestDeleteUser", t, func() {
		w := httptest.NewRecorder()
		cookie := getAdminCookie(r)
		req, _ := http.NewRequest("DELETE", "/api/user/2", nil)
		req.Header.Set("Cookie", cookie)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 200)

		w = httptest.NewRecorder()
		req, _ = http.NewRequest("GET", "/api/user/2", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 404)
	})

}
func TestUserChangePassword(t *testing.T) {
	restoreTestDatabase()
	r := GetGinEngine()
	Convey("TestUserChangePassword", t, func() {
		w := httptest.NewRecorder()
		payload := map[string]string{
			"oldPassword": "123456",
			"newPassword": "123456789",
		}
		data, _ := json.Marshal(payload)
		req, _ := http.NewRequest("PUT", "/api/user/1/password", bytes.NewBuffer(data))
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 200)

		w = httptest.NewRecorder()
		req, _ = http.NewRequest("GET", "/api/user/1", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 200)
		body, err := io.ReadAll(w.Body)
		So(err, ShouldBeNil)

		var res model.User
		err = json.Unmarshal(body, &res)
		So(err, ShouldBeNil)

		So(res.UserName, ShouldEqual, "admin")
		So(res.IsAdmin, ShouldEqual, true)
		So(res.VerifyPassword("123456789"), ShouldEqual, true)
	})
}
