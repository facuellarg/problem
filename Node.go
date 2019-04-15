package problem

import (
	"bytes"
)

//Node struc
type Node struct {
	p    Problem
	path bytes.Buffer
}

//NewNode return new node
func NewNode(p Problem, path string) (this Node) {
	this.p = p
	this.path.WriteString(path)
	return
}

//Path return a path
func (n Node) Path() string {
	return n.path.String()
}

//Problem return a problem
func (n Node) Problem() Problem {
	return n.p
}

//AddPath add a path
func (n *Node) AddPath(path string) {
	n.path.WriteString(path)
}
