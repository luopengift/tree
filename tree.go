package tree

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/luopengift/log"
)

// Tree tree
type Tree struct {
	root  string
	Delim string
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
func InitTree(root, delim string) *Tree {
	return &Tree{
		Node:  InitNode(root),
		Delim: delim,
	}
}

// Root root
func (tr *Tree) Root() *Node {
	return tr.Node
}

// Get get
func (tr *Tree) Get(name string) *Node {
	current := tr.Root()
	keys := strings.Split(name, tr.Delim)
	for _, key := range keys {
		if key == "" {
			continue
		}
		child := current.Get(key)
		if child == nil {
			return nil
		}
		current = child
	}
	return current
}

// Put put
func (tr *Tree) Put(name string, value ...interface{}) {
	keys := strings.Split(name, tr.Delim)
	current := tr.Root()
	log.Info("%v, %#v", name, keys)
	for _, key := range keys {
		if key == "" {
			continue
		}
		child := current.Get(key)
		if child == nil {
			child = InitNode(key)
			current.AddChild(child)
		}
		current = child
	}
	if len(value) != 0 {
		current.SetValue(value[0])
	}
}

// Delete delete a branch or node
func (tr *Tree) Delete(name string) {
	keys := strings.Split(name, tr.Delim)
	current := tr.Root()
	for _, key := range keys {
		if key == "" {
			continue
		}
		child := current.Get(key)
		if child == nil {
			return
		}
		current = child
	}
	current.parent.Delete(current.Name)
}

// Load load yaml file
func (tr *Tree) Load(filepath string) error {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, tr)
}

// Dump dump to file
func (tr *Tree) Dump(filepath string) error {
	b, err := json.MarshalIndent(tr, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath, b, 0644)
}
