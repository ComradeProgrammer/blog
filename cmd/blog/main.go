package main

import (
	"log"

	"github.com/ComradeProgrammer/blog/internal/blog/controller"
	"github.com/ComradeProgrammer/blog/internal/blog/model"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	model.ConnectDatabase("database.sqlite")
	model.InitDatabase()
	controller.GetGinEngine().Run(":80")
}
