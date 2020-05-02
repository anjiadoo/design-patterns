package node

import (
	"design-patterns/component/demo1/component"
	"net/http"
)

type Node struct {
	Name        string
	Url         string
	NodeHandler component.NodeHandler
}

func NewNode(name string, handler component.NodeHandler) *Node {
	return &Node{Name: name, NodeHandler: handler}
}

func (pN *Node) Add(iDirComponent component.DirComponent) bool {
	return false
}

func (pN *Node) Remove(name string) bool {
	return false
}

func (pN *Node) GetChild() (sziDirComponent []component.DirComponent, b bool) {
	return
}

func (pN *Node) Handler(w http.ResponseWriter, req *http.Request) {
	pN.NodeHandler(w, req)
}

func (pN *Node) GetName() string {
	return pN.Name
}

func (pN *Node) IsDir() bool {
	return false
}
