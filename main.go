package main

import (
	"fmt"
)

func makeTree() *Node {
	root := newNode(1)
	firstOne := newNode(2)
	firstTwo := newNode(3)
	firstThree := newNode(4)
	secondOne := newNode(5)
	secondTwo := newNode(6)
	secondThree := newNode(7)
	root.append(firstOne)
	root.append(firstTwo)
	root.append(firstThree)
	firstOne.append(secondOne)
	firstOne.append(secondTwo)
	secondThree.append(secondThree)
	return root
}

func main() {
	root := makeTree()
	root.printChildren()
	fmt.Println("HI")
}
