package tree

import (
	"fmt"
	"strings"
)

// Tree tree
type Tree struct {
	root string
	*Node
}

// GOString go string
func (tr *Tree) GOString() string {
	var root = tr.Name
	for _, node := range tr.Next {
		fmt.Print(node.Name)
	}

	return root
}

// InitTree init tree
func InitTree() *Tree {
	return &Tree{
		Node: InitNode("/"),
	}
}

// Root root
func (tr *Tree) Root() *Node {
	return tr.Node
}

// Get get
func (tr *Tree) Get(name string) *Node {
	return nil
}

// Put put
func (tr *Tree) Put(name string) {
	//split := tr.Root().Name
	keys := strings.Split(name[1:], "/")
	node := tr.Root()
	for _, key := range keys {
		if child := node.Get(key); child == nil {
			child = InitNode(key)
			node.Set(child)
			node = child
		}
	}
	fmt.Println(keys)

}

// Load load yaml file
func (tr *Tree) Load(file string) error {
	return nil
}
