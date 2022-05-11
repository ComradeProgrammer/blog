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

func TestGetBlogs(t *testing.T) {
	restoreTestDatabase()
	r := GetGinEngine()
	Convey("TestGetBlogs", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/blog", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 200)
		body, err := io.ReadAll(w.Body)
		So(err, ShouldBeNil)

		var res []model.Blog
		err = json.Unmarshal(body, &res)
		So(err, ShouldBeNil)
		So(len(res), ShouldEqual, 4)
		So(res[0].Title, ShouldEqual, "title4")
		So(res[0].Content, ShouldEqual, "content4")
		So(res[0].CategoryID, ShouldEqual, 2)
		So(res[1].Title, ShouldEqual, "title3")
		So(res[1].Content, ShouldEqual, "content3")
		So(res[1].CategoryID, ShouldEqual, 2)
		So(res[2].Title, ShouldEqual, "title2")
		So(res[2].Content, ShouldEqual, "content2")
		So(res[2].CategoryID, ShouldEqual, 1)
		So(res[3].Title, ShouldEqual, "title1")
		So(res[3].Content, ShouldEqual, "content1")
		So(res[3].CategoryID, ShouldEqual, 1)
	})
}

func TestGetBlog(t *testing.T) {
	restoreTestDatabase()
	r := GetGinEngine()
	Convey("TestGetBlog", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/blog/2", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 200)
		body, err := io.ReadAll(w.Body)
		So(err, ShouldBeNil)

		var res model.Blog
		err = json.Unmarshal(body, &res)
		So(err, ShouldBeNil)
		So(res.Title, ShouldEqual, "title2")
		So(res.Content, ShouldEqual, "content2")
		So(res.CategoryID, ShouldEqual, 1)
		So(res.Category.ID, ShouldEqual, 1)
	})
}
func TestPostBlog(t *testing.T) {
	restoreTestDatabase()
	r := GetGinEngine()
	Convey("TestPostBlogUnauthorized", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/blog", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 401)
	})

	Convey("TestPostBlog", t, func() {
		w := httptest.NewRecorder()
		cookie := getAdminCookie(r)
		blog := model.Blog{
			Title:      "MyBlog",
			Content:    "BlogContent",
			CategoryID: 2,
		}
		data, _ := json.Marshal(blog)
		req, _ := http.NewRequest("POST", "/api/blog", bytes.NewBuffer(data))
		req.Header.Set("Cookie", cookie)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 200)

		w = httptest.NewRecorder()
		req, _ = http.NewRequest("GET", "/api/blog/5", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 200)
		body, err := io.ReadAll(w.Body)
		So(err, ShouldBeNil)
		var res model.Blog
		err = json.Unmarshal(body, &res)
		So(err, ShouldBeNil)
		So(res.Title, ShouldEqual, blog.Title)
		So(res.Content, ShouldEqual, blog.Content)
		So(res.Category.ID, ShouldEqual, 2)
	})

	Convey("TestPostBlogInvalid", t, func() {
		w := httptest.NewRecorder()
		cookie := getAdminCookie(r)
		blog := model.Blog{
			Title:   "MyBlog",
			Content: "BlogContent",
		}
		data, _ := json.Marshal(blog)
		req, _ := http.NewRequest("POST", "/api/blog", bytes.NewBuffer(data))
		req.Header.Set("Cookie", cookie)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 400)

	})

	Convey("TestPostBlogInvalid", t, func() {
		w := httptest.NewRecorder()
		cookie := getAdminCookie(r)
		blog := model.Blog{
			Title:      "MyBlog",
			Content:    "BlogContent",
			CategoryID: 0,
		}
		data, _ := json.Marshal(blog)
		req, _ := http.NewRequest("POST", "/api/blog", bytes.NewBuffer(data))
		req.Header.Set("Cookie", cookie)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 400)
	})

	Convey("TestDeleteBlog", t, func() {
		w := httptest.NewRecorder()
		cookie := getAdminCookie(r)
		req, _ := http.NewRequest("DELETE", "/api/blog/3", nil)
		req.Header.Set("Cookie", cookie)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 200)

		w = httptest.NewRecorder()
		req, _ = http.NewRequest("GET", "/api/blog/3", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 404)
	})
}

func TestPutBlog(t *testing.T) {
	restoreTestDatabase()
	r := GetGinEngine()
	Convey("TestPutBlogUnauthorized", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/api/blog/1", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 401)
	})

	Convey("TestPutBlog", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/blog/2", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 200)
		body, err := io.ReadAll(w.Body)
		So(err, ShouldBeNil)

		var oldBlog model.Blog
		err = json.Unmarshal(body, &oldBlog)

		w = httptest.NewRecorder()
		cookie := getAdminCookie(r)
		modifiedBlog := model.Blog{
			Title:      "modified",
			Content:    "modifiedDescription",
			CategoryID: 1,
		}
		data, err := json.Marshal(modifiedBlog)
		req2, _ := http.NewRequest("PUT", "/api/blog/2", bytes.NewBuffer(data))
		req2.Header.Set("Cookie", cookie)
		r.ServeHTTP(w, req2)
		So(w.Code, ShouldEqual, 200)

		w = httptest.NewRecorder()
		req3, _ := http.NewRequest("GET", "/api/blog/2", nil)
		r.ServeHTTP(w, req3)
		So(w.Code, ShouldEqual, 200)
		body, err = io.ReadAll(w.Body)
		So(err, ShouldBeNil)

		var newBlog model.Blog
		err = json.Unmarshal(body, &newBlog)

		So(newBlog.Title, ShouldEqual, modifiedBlog.Title)
		So(newBlog.Content, ShouldEqual, modifiedBlog.Content)
		So(newBlog.CategoryID, ShouldEqual, newBlog.CategoryID)
		So(newBlog.CreateAt.Equal(oldBlog.CreateAt), ShouldBeTrue)
	})
}

func TestDeleteBlog(t *testing.T) {
	restoreTestDatabase()
	r := GetGinEngine()

	Convey("TestDeleteBlogUnauthorized", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/api/blog/1", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 401)
	})
	Convey("TestDeleteBlog", t, func() {
		w := httptest.NewRecorder()
		cookie := getAdminCookie(r)
		req, _ := http.NewRequest("DELETE", "/api/blog/300", nil)
		req.Header.Set("Cookie", cookie)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 404)
	})

}
