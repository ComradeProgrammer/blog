package main

import (
	"log"

	"github.com/ComradeProgrammer/blog/internal/myblog/controller"
	"github.com/ComradeProgrammer/blog/internal/myblog/dal/conn"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	db, err := conn.ConnectSqliteDatabase("database.sqlite")
	if err != nil {
		panic(err)
	}

	conn.DB = db

	r, err := controller.GetGinEngine()
	if err != nil {
		panic(err)
	}
	r.Run(":8080")
	// r.Use(TlsHandler())
	// r.RunTLS(":443", "zhizhangertong.top.crt", "zhizhangertong.top.key")
}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
