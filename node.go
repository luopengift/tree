package tree

import (
	"sync"
)

// Node node
type Node struct {
	Name    string
	parent  *Node
	self    *Node // This is debug
	Next    []*Node
	Value   interface{}
	version int
	mux     *sync.RWMutex
}

// InitNode init node
func InitNode(name string) *Node {
	node := &Node{
		Name: name,
	}
	node.self = node
	return node
}

// AddChild set, 增加子节点
func (n *Node) AddChild(nodes ...*Node) {
	n.Next = append(n.Next, nodes...)
	for _, node := range nodes {
		node.parent = n
	}
}

// SetValue set value to node
func (n *Node) SetValue(v interface{}) {
	n.Value = v
}

// Get get , 查询当前这一层
func (n *Node) Get(name string) *Node {
	for _, node := range n.Next {
		if node.Name == name {
			return node
		}
	}
	return nil
}

// Delete a node and all subnode
func (n *Node) Delete(name string) *Node {
	//node := n.Search(name)

	return nil
}

// IsExist is exist,是否存在
func (n *Node) IsExist(name string) bool {
	return n.Get(name) != nil
}

// Search , 遍历所有子节点
func (n *Node) Search(name string) *Node {
	for _, node := range n.Next {
		if node.Name == name {
			return node
		}
		return node.Search(name)
	}
	return nil
}
