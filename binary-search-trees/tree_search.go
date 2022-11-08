package binarysearchtrees

func TreeSearch(treeNode *Node[int32], k int32) *Node[int32] {
	if treeNode == nil || treeNode.Value == k {
		return treeNode
	}
	if k < treeNode.Value {
		return TreeSearch(treeNode.Left, k)
	} else {
		return TreeSearch(treeNode.Right, k)
	}
}

func IterativeSearch(treeNode *Node[int32], k int32) *Node[int32] {
	head := treeNode
	for head != nil {
		if head.Value == k {
			return head
		}
		if k < head.Value {
			head = head.Left
		} else {
			head = head.Right
		}
	}
	return head
}

func Minimum(treeNode *Node[int32]) *Node[int32] {
	for treeNode.Left != nil {
		treeNode = treeNode.Left
	}
	return treeNode
}

func Maximum(treeNode *Node[int32]) *Node[int32] {
	for treeNode.Right != nil {
		treeNode = treeNode.Right
	}
	return treeNode
}

func MinimumRecursive(treeNode *Node[int32]) *Node[int32] {
	if treeNode.Left == nil {
		return treeNode
	}
	return MinimumRecursive(treeNode.Left)
}

func MaximumRecursive(treeNode *Node[int32]) *Node[int32] {
	if treeNode.Right == nil {
		return treeNode
	}
	return MaximumRecursive(treeNode.Right)
}

func TreeSuccessor(node *Node[int32]) *Node[int32] {
	if node.Right != nil {
		return Minimum(node.Right)
	} else {
		parent := node.Parent
		for parent != nil && parent.Right == node {
			node = parent
			parent = node.Parent
		}
		return parent
	}
}

func TreePreDecessor(node *Node[int32]) *Node[int32] {
	if node.Left != nil {
		return Maximum(node.Left)
	} else {
		parent := node.Parent
		for parent != nil && node == parent.Left {
			node = parent
			parent = node.Parent
		}
		return parent
	}
}
