package api

import (
	"github.com/luopengift/gohttp"
	"github.com/luopengift/tree"
)

var tr = tree.InitTree("/", "/")

// Init api
func Init() {
	app := gohttp.Init()
	app.RouteFunCtx("^/_dump$", func(ctx *gohttp.Context) {
		if err := tr.Dump("txt.txt"); err != nil {
			ctx.Output(err)
			return
		}
		ctx.Output("dump ok!")
	})
	app.RouteFunCtx("^/_load$", func(ctx *gohttp.Context) {
		if err := tr.Load("txt.txt"); err != nil {
			ctx.Output(err)
			return
		}
		ctx.Output("load ok!")
	})
	app.Route("^/.*$", &TreeHandler{})

	app.Run(":8080")
}

// TreeHandler TreeHandler
type TreeHandler struct {
	gohttp.APIHandler
}

// GET get
func (t *TreeHandler) GET() {
	node := tr.Get(t.URL.String())
	if node == nil {
		t.Output("path not exist")
		return
	}
	t.Output(node.Value)
}

// POST post
func (t *TreeHandler) POST() {
	tr.Put(t.URL.String(), t.GetBodyArgs())
	t.Output("ok")
}

// DELETE delete
func (t *TreeHandler) DELETE() {
	tr.Delete(t.URL.String())
}
