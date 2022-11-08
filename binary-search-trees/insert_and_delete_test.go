package binarysearchtrees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	node := &Node[int32]{Value: 7}
	tree := BuildTree()
	fiveNode := TreeSearch(tree, 5)
	assert.Nil(t, fiveNode.Right)
	assert.Nil(t, fiveNode.Left)
	TreeInsert(tree, node)
	assert.NotNil(t, fiveNode.Right)
	assert.Nil(t, fiveNode.Left)
	assert.Equal(t, int32(7), fiveNode.Right.Value)
}
