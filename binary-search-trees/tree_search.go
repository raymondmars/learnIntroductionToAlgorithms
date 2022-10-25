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
