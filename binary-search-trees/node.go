package binarysearchtrees

type Node[T string | int32] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

// list := []int{1, 4, 5, 10, 16, 17, 21}

// func BuildTree(node *Node) *Node {
// 	return nil
// }
