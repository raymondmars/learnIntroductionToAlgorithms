package binarysearchtrees

import "fmt"

func InorderTreeWalk(node *Node) {
	if node != nil {
		InorderTreeWalk(node.Left)
		fmt.Println(node.Value)
		InorderTreeWalk(node.Right)
	}
}
