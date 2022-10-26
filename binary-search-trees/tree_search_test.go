package binarysearchtrees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeSearch(t *testing.T) {
	node := TreeSearch(BuildTree(), 17)
	assert.NotNil(t, node)
	assert.Equal(t, int32(17), node.Value)

	node1 := TreeSearch(BuildTree(), 170)
	assert.Nil(t, node1)
}

func TestInterativeSearch(t *testing.T) {
	node := IterativeSearch(BuildTree(), 17)
	assert.NotNil(t, node)
	assert.Equal(t, int32(17), node.Value)

	node1 := IterativeSearch(BuildTree(), 170)
	assert.Nil(t, node1)
}

func TestMinimumSearch(t *testing.T) {
	assert.Equal(t, int32(1), Minimum(BuildTree()).Value)
}

func TestMaximumSearch(t *testing.T) {
	assert.Equal(t, int32(21), Maximum(BuildTree()).Value)
}

func TestMinimumRecursiveSearch(t *testing.T) {
	assert.Equal(t, int32(1), MinimumRecursive(BuildTree()).Value)
}

func TestMaximumResursiveSearch(t *testing.T) {
	assert.Equal(t, int32(21), MaximumRecursive(BuildTree()).Value)
}

func TestSuccessor(t *testing.T) {
	tree := BuildTree()
	fiveNode := TreeSearch(tree, 5)
	assert.Equal(t, int32(10), TreeSuccessor(fiveNode).Value)

	fourNode := TreeSearch(tree, 4)
	assert.Equal(t, int32(5), TreeSuccessor(fourNode).Value)
}

func TestPreDecessor(t *testing.T) {
	tree := BuildTree()
	fiveNode := TreeSearch(tree, 5)
	assert.Equal(t, int32(4), TreePreDecessor(fiveNode).Value)

	tenNode := TreeSearch(tree, 10)
	assert.Equal(t, int32(5), TreePreDecessor(tenNode).Value)

	sixteenNode := TreeSearch(tree, 16)
	assert.Equal(t, int32(10), TreePreDecessor(sixteenNode).Value)
}
