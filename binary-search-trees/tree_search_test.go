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
