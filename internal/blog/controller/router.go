package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func GetGinEngine() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session", store))
	r.StaticFS("/api/static", http.Dir("static"))
	r.Use(static.Serve("/", static.LocalFile("web/build", true)))

	//keep-alive handler
	r.GET("/api/ping", Ping)

	//user handler
	r.POST("/api/login", Login)
	r.DELETE("/api/login", LogOut)

	//blog catrgory handler
	r.GET("/api/category/:id", GetCategory)
	r.GET("/api/category", GetCategories)
	r.POST("/api/category", PostCategory)
	r.PUT("/api/category/:id", PutCategory)
	r.DELETE("/api/category/:id", DeleteCategory)

	//blog handler
	r.GET("/api/blog", GetBlogs)
	r.GET("/api/blog/:id", GetBlog)
	r.POST("/api/blog", PostBlog)
	r.PUT("/api/blog/:id", PutBlog)
	r.DELETE("/api/blog/:id", DeleteBlog)

	//comment handler
	r.POST("/api/comment", PostComment)
	r.DELETE("/api/comment/:id", DeleteComment)

	//user controller
	r.GET("/api/user", GetUsers)
	r.GET("/api/user/:id", GetUser)
	r.POST("/api/user", PostUser)
	r.DELETE("/api/user/:id", DeleteUser)

	r.PUT("/api/user/:username/password", PutUserPassword)
	return r
}
