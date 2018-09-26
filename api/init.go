package api

import (
	"github.com/luopengift/gohttp"
	"github.com/luopengift/log"
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

type TreeHandler struct {
	gohttp.APIHandler
}

func (t *TreeHandler) Prepare() {
	log.Info("prepare, %v", t.URL)
}

func (t *TreeHandler) GET() {
	node := tr.Get(t.URL.String())
	if node == nil {
		t.Output("path not exist")
		return
	}
	t.Output(node.Value)
}

func (t *TreeHandler) POST() {
	tr.Put(t.URL.String(), t.GetBodyArgs())

	t.Output("ok")
}
