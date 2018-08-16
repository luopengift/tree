package main

import (
	"fmt"

	"github.com/luopengift/log"
	"github.com/luopengift/tree"
)

func main() {

	fmt.Println("init")
	tr := tree.InitTree()
	fmt.Println(tr.Node)
	fmt.Println(tr.Node)
	tr.Put("/test/config/luopeng")
	fmt.Println(tr.Node)
	fmt.Println("=====")
	for k, v := range tr.Node.Next {
		fmt.Println(k, v)
	}
	log.Display("$$", tr.Node)
}
