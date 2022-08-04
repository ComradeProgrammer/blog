package main

import (
	"log"

	"github.com/ComradeProgrammer/blog/internal/blog/controller"
	"github.com/ComradeProgrammer/blog/internal/blog/model"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	model.ConnectDatabase("database.sqlite")
	model.InitDatabase()
	r := controller.GetGinEngine()

	//r.Run(":8080")
	r.Use(TlsHandler())
	r.RunTLS(":443", "zhizhangertong.top.crt", "zhizhangertong.top.key")
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
