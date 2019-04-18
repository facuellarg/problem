package problem

import (
	"bytes"
)

//Node struc
type Node struct {
	p        Problem
	path     bytes.Buffer
	pathCost int
	rango    int
	index    int
}

//NewNode return new node
func NewNode(p Problem, path string, cost int) (this Node) {
	this.p = p
	this.path.WriteString(path)
	this.pathCost = cost
	return
}

//Path return a path
func (n Node) Path() string {
	return n.path.String()
}

//Problem return a problem.
func (n Node) Problem() Problem {
	return n.p
}

//Cost return the current cost until the state.
func (n Node) Cost() int { return n.pathCost }

//SetRange set the range into the node
func (n *Node) SetRange(r int) { n.rango = r }

//AddPath add a path.
// func (n *Node) AddPath(path string) {
// 	n.path.WriteString(path)
// }
