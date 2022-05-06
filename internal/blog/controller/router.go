package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func GetGinEngine() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session", store))

	//keep-alive handler
	r.GET("/ping", Ping)

	//user handler
	r.POST("/login", Login)
	r.DELETE("login", LogOut)

	//blog catrgory handler
	r.GET("/category/:id", GetCategory)
	r.GET("/category", GetCategories)
	r.POST("/category", PostCategory)
	r.PUT("/category/:id", PutCategory)
	r.DELETE("/category/:id", DeleteCategory)

	//blog handler
	r.GET("/blog",GetBlogs)
	r.GET("/blog/:id",GetBlog)
	r.POST("/blog",PostBlog)
	r.PUT("/blog/:id",PutBlog)
	r.DELETE("/blog/:id",DeleteBlog)

	//user controller
	r.GET("/user",GetUsers)
	r.GET("/user/:id",GetUser)
	r.POST("/user",PostUser)
	r.DELETE("/user/:id",DeleteUser)
	
	r.PUT("/user/:id/password",PutUserPassword)
	return r
}
