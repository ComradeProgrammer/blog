package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPing(t *testing.T) {
	restoreTestDatabase()
	r := GetGinEngine()
	Convey("TestPingBasic", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/ping", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 200)
	})
}
