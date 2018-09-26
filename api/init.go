package api

import (
	"github.com/luopengift/gohttp"
	"github.com/luopengift/tree"
)

var tr = tree.InitTree("/", "/")

// Init api
func Init() {
	app := gohttp.Init()
	app.Run()
}
