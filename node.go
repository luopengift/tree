package tree

import (
	"sync"
)

// Node node
type Node struct {
	Name   string
	parent *Node
	Next   []*Node
	Value  interface{}
	mux    *sync.RWMutex
}

// InitNode init node
func InitNode(name string) *Node {
	return &Node{
		Name: name,
	}
}

// Set set
func (n *Node) Set(node *Node) {
	n.Next = append(n.Next, node)
	node.parent = n
}

// Get get
func (n *Node) Get(name string) *Node {
	if n.Name == name {
		return n
	}
	for _, node := range n.Next {
		if node.Name == name {
			return node
		}
		return node.Get(name)
	}
	return nil
}
