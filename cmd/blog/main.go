package main

import (
	"github.com/ComradeProgrammer/blog/internal/blog/model"
	"github.com/ComradeProgrammer/blog/internal/blog/router"
)

func main() {
	model.InitDatabase("database.sqlite")
	router.GetGinEngine().Run()
}
