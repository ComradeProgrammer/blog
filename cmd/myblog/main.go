package main

import (
	"flag"
	"log"

	"github.com/ComradeProgrammer/blog/internal/myblog/controller"
	"github.com/ComradeProgrammer/blog/internal/myblog/dal/conn"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

var sqliteDatabase = flag.String("sqlitedb", "database.sqlite", "database file for the sqlite database")
var sslCertPath = flag.String("ssl-cert", "", "ssl certification path")
var sslKeyPath = flag.String("ssl-key", "", "ssl certification path")

func main() {

	flag.Parse()
	//gin.SetMode(gin.ReleaseMode)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	db, err := conn.ConnectSqliteDatabase(*sqliteDatabase)
	if err != nil {
		panic(err)
	}

	conn.DB = db

	r, err := controller.GetGinEngine()
	if err != nil {
		panic(err)
	}
	if *sslCertPath == "" || *sslKeyPath == "" {
		r.Run(":8080")
	} else {
		r.Use(TlsHandler())
		r.RunTLS(":443", *sslCertPath, *sslKeyPath)
	}
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
