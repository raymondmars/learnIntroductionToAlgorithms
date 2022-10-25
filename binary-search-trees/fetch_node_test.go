package binarysearchtrees

import "testing"

func TestInOrderTreeWalk(t *testing.T) {
	trees := BuildTree()
	InorderTreeWalk(trees)
}
