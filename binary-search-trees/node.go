package binarysearchtrees

type Node[T string | int32] struct {
	Value  T
	Parent *Node[T]
	Left   *Node[T]
	Right  *Node[T]
}

func BuildTree() *Node[int32] {
	// list := []int{1, 4, 5, 10, 16, 17, 21}
	node1 := &Node[int32]{Value: 10}
	node2 := &Node[int32]{Value: 4}
	node3 := &Node[int32]{Value: 17}
	node4 := &Node[int32]{Value: 1}
	node5 := &Node[int32]{Value: 5}
	node6 := &Node[int32]{Value: 16}
	node7 := &Node[int32]{Value: 21}

	node1.Left = node2
	node1.Right = node3
	node1.Parent = nil

	node2.Parent = node1
	node2.Left = node4
	node2.Right = node5

	node3.Parent = node1
	node3.Left = node6
	node3.Right = node7

	node4.Parent = node2
	node5.Parent = node2

	node6.Parent = node3
	node7.Parent = node3

	return node1
}
