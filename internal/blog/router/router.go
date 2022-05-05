package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/ComradeProgrammer/blog/internal/blog/controller"
)

func GetGinEngine() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session", store))

	//keep-alive handler
	r.GET("/ping", controller.Ping)

	//user handler
	r.POST("/login", controller.Login)
	r.DELETE("login", controller.LogOut)

	//blog catrgory handler
	r.GET("/categories",controller.GetCategories)
	r.GET("/category/:id",controller.GetCategory)
	r.POST("/category",controller.PostCategory)
	r.PUT("/category/:id",controller.PutCategory)
	r.DELETE("/category/:id",controller.DeleteCategory)

	return r
}
