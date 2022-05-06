package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func getAdminCookie(r *gin.Engine) string {
	w := httptest.NewRecorder()
	var body = map[string]string{
		"userName": "admin",
		"password": "123456",
	}
	data, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(data))
	r.ServeHTTP(w, req)
	return w.Header().Get("Set-Cookie")

}
