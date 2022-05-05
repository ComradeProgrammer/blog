package main

import (
	"log"

	"github.com/ComradeProgrammer/blog/internal/blog/model"
	"github.com/ComradeProgrammer/blog/internal/blog/router"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	model.ConnectDatabase("database.sqlite")
	model.InitDatabase()
	router.GetGinEngine().Run()
}
