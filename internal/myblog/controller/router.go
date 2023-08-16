package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetGinEngine(db *gorm.DB)( *gin.Engine,error ){
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session", store))
	r.StaticFS("/api/static", http.Dir("static"))
	r.Use(static.Serve("/",static.LocalFile("web/build",true)))


	pingController:=NewPingController()
	blogController,err:=NewBlogController(db)
	if err!=nil{
		return nil,err
	}
	categoryController,err:=NewCategoryController(db)
	if err!=nil{
		return nil,err
	}
	loginController,err:=NewLoginController(db)
	if err!=nil{
		return nil,err
	}


	//keep-alive handler
	r.GET("/api/ping", pingController.Ping)

	//user handler
	r.POST("/api/login", loginController.Login)
	r.DELETE("/api/login", loginController.LogOut)

	//blog catrgory handler
	r.GET("/api/category/:id", categoryController.GetCategory)
	r.GET("/api/category", categoryController.ListCategory)
	r.POST("/api/category", categoryController.PostCategory)
	r.PUT("/api/category/:id", categoryController.PutCategory)
	r.DELETE("/api/category/:id", categoryController.DeleteCategory)

	//blog handler
	r.GET("/api/blog", blogController.ListBlogs)
	r.GET("/api/blog/:id", blogController.GetBlog)
	r.POST("/api/blog", blogController.PostBlog)
	r.PUT("/api/blog/:id", blogController.PutBlog)
	r.DELETE("/api/blog/:id", blogController.DeleteBlog)

	//comment handler
	// r.POST("/api/comment", PostComment)
	// r.DELETE("/api/comment/:id",DeleteComment)

	//user controller
	// r.GET("/api/user", GetUsers)
	// r.GET("/api/user/:id", GetUser)
	// r.POST("/api/user", PostUser)
	// r.DELETE("/api/user/:id", DeleteUser)

	// r.PUT("/api/user/:username/password", PutUserPassword)
	return r,nil
}
