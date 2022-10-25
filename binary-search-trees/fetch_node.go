package binarysearchtrees

import "fmt"

func InorderTreeWalk(node *Node[int32]) {
	if node != nil {
		InorderTreeWalk(node.Left)
		fmt.Println(node.Value)
		InorderTreeWalk(node.Right)
	}
}

func PreOrderTreeWalk(node *Node[int32]) {
	if node != nil {
		fmt.Println(node.Value)
		PreOrderTreeWalk(node.Left)
		PreOrderTreeWalk(node.Right)
	}
}

func PostOrderTreeWalk(node *Node[int32]) {
	if node != nil {
		PostOrderTreeWalk(node.Left)
		PostOrderTreeWalk(node.Right)
		fmt.Println(node.Value)
	}
}
