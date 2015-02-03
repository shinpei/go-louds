package main

import (
	"fmt"
)

type Node struct {
	value    int64
	children []*Node
}

func newNode(val int64) *Node {
	return &Node{value: val, children: make([]*Node, 0)}
}

func (self *Node) append(node *Node) {
	self.children = append(self.children, node)
}

func (self *Node) printChildren() {
	fmt.Println("value=", self.value)
	for _, n := range self.children {
		n.printChildren()
	}
}
