package binarysearchtrees

func TreeInsert(tree *Node[int32], node *Node[int32]) *Node[int32] {
	if tree == nil {
		return node
	}

	var parent *Node[int32]
	searchPointer := tree
	for searchPointer != nil {
		parent = searchPointer
		if node.Value > searchPointer.Value {
			searchPointer = searchPointer.Right
		} else {
			searchPointer = searchPointer.Left
		}
	}
	if parent != nil {
		node.Parent = parent
		if node.Value > parent.Value {
			parent.Right = node
		} else {
			parent.Left = node
		}
	}
	return tree
}

func TreeDelete(tree *Node[int32], node *Node[int32]) {

}
