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

func TestGetCategories(t *testing.T) {
	restoreTestDatabase()
	r := GetGinEngine()
	Convey("TestGetCategoriesBasic", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/category", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 200)
		body, err := io.ReadAll(w.Body)
		So(err, ShouldBeNil)
		var res []model.Category
		err = json.Unmarshal(body, &res)
		So(err, ShouldBeNil)

		So(len(res), ShouldEqual, 2)
		So(res[0].Name, ShouldEqual, "category2")
		So(res[0].Description, ShouldEqual, "description2")

		So(res[1].Name, ShouldEqual, "category1")
		So(res[1].Description, ShouldEqual, "description1")

	})

}

func TestGetCategory(t *testing.T) {

	restoreTestDatabase()
	r := GetGinEngine()
	Convey("TestGetCategory", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/category/1", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 200)
		body, err := io.ReadAll(w.Body)
		So(err, ShouldBeNil)

		var res model.Category
		err = json.Unmarshal(body, &res)
		So(err, ShouldBeNil)

		So(res.Name, ShouldEqual, "category1")
		So(res.Description, ShouldEqual, "description1")
		So(len(*res.Blogs), ShouldEqual, 2)

		So((*res.Blogs)[0].Title, ShouldEqual, "title2")
		So((*res.Blogs)[0].Content, ShouldEqual, "content2")
		So((*res.Blogs)[1].Title, ShouldEqual, "title1")
		So((*res.Blogs)[1].Content, ShouldEqual, "content1")
	})
	Convey("TestGetCategory404", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/category/100", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 404)
	})
}

func TestPostCategory(t *testing.T) {
	restoreTestDatabase()
	r := GetGinEngine()
	Convey("TestPostCategoryUnauthorized", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/category", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 401)
	})
	Convey("TestGetCategory", t, func() {
		w := httptest.NewRecorder()
		cookie := getAdminCookie(r)

		category := model.Category{
			Name:        "category3",
			Description: "description3",
		}
		data, _ := json.Marshal(category)

		req, _ := http.NewRequest("POST", "/category", bytes.NewBuffer(data))
		req.Header.Set("Cookie", cookie)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 200)

		w = httptest.NewRecorder()
		req, _ = http.NewRequest("GET", "/category/3", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 200)
		body, err := io.ReadAll(w.Body)
		So(err, ShouldBeNil)

		var res model.Category
		err = json.Unmarshal(body, &res)
		So(err, ShouldBeNil)

		So(res.Name, ShouldEqual, "category3")
		So(res.Description, ShouldEqual, "description3")
		So(len(*res.Blogs), ShouldEqual, 0)
	})
	Convey("TestDeleteCategory", t, func() {
		w := httptest.NewRecorder()
		cookie := getAdminCookie(r)
		req, _ := http.NewRequest("DELETE", "/category/3", nil)
		req.Header.Set("Cookie", cookie)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 200)

		w = httptest.NewRecorder()
		req, _ = http.NewRequest("GET", "/category/3", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 404)
	})

}
func TestPutCategory(t *testing.T) {
	restoreTestDatabase()
	r := GetGinEngine()
	Convey("TestPutCategoryUnauthorized", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/category/1", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 401)
	})
	Convey("TestPutCategory", t, func() {
		w := httptest.NewRecorder()
		cookie := getAdminCookie(r)
		modifiedCategory := model.Category{
			Name:        "modified",
			Description: "modifiedDescription",
		}
		data, _ := json.Marshal(modifiedCategory)
		req2, _ := http.NewRequest("PUT", "/category/200", bytes.NewBuffer(data))
		req2.Header.Set("Cookie", cookie)
		r.ServeHTTP(w, req2)
		So(w.Code, ShouldEqual, 404)
	})
	Convey("TestPutCategory", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/category/2", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 200)
		body, err := io.ReadAll(w.Body)
		So(err, ShouldBeNil)

		var oldCategory model.Category
		err = json.Unmarshal(body, &oldCategory)

		w = httptest.NewRecorder()
		cookie := getAdminCookie(r)

		modifiedCategory := model.Category{
			Name:        "modified",
			Description: "modifiedDescription",
		}
		data, err := json.Marshal(modifiedCategory)
		req2, _ := http.NewRequest("PUT", "/category/2", bytes.NewBuffer(data))
		req2.Header.Set("Cookie", cookie)
		r.ServeHTTP(w, req2)
		So(w.Code, ShouldEqual, 200)

		w = httptest.NewRecorder()
		req3, _ := http.NewRequest("GET", "/category/2", nil)
		r.ServeHTTP(w, req3)
		So(w.Code, ShouldEqual, 200)
		body, err = io.ReadAll(w.Body)
		So(err, ShouldBeNil)

		var newCategory model.Category
		err = json.Unmarshal(body, &newCategory)

		So(newCategory.ID, ShouldEqual, 2)
		So(newCategory.Name, ShouldEqual, modifiedCategory.Name)
		So(newCategory.Description, ShouldEqual, modifiedCategory.Description)

		So(newCategory.CreateAt.Equal(oldCategory.CreateAt), ShouldBeTrue)

	})
}

func TestDeleteCategory(t *testing.T) {
	restoreTestDatabase()
	r := GetGinEngine()

	Convey("TestDeleteCategoryUnauthorized", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/category/1", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 401)
	})
	Convey("TestDeleteCategory", t, func() {
		w := httptest.NewRecorder()
		cookie := getAdminCookie(r)
		req, _ := http.NewRequest("DELETE", "/category/300", nil)
		req.Header.Set("Cookie", cookie)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 404)
	})

}
