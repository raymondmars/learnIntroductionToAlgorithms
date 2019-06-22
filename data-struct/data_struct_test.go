package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func Equal(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{})

//Test all queue features
func TestQueue(t *testing.T) {
	queue := &Queue{}
	assert.True(t, queue.IsEmpty())
	queue.Push(1)
	assert.False(t, queue.IsEmpty())
	queue.Push(2)
	assert.Equal(t, 2, queue.Len())

	assert.Equal(t, 1, queue.Top())
	assert.Equal(t, 1, queue.Pop())
	assert.Equal(t, 2, queue.Top())

	queue.Pop()
	assert.Equal(t, nil, queue.Pop())
	assert.Equal(t, 0, queue.Len())
}

//Test all stack features
func TestStack(t *testing.T) {
	stack := &Stack{}
	assert.True(t, stack.IsEmpty())
	stack.Push(1)
	assert.False(t, stack.IsEmpty())
	stack.Push(2)
	assert.Equal(t, 2, stack.Len())

	assert.Equal(t, 2, stack.Top())
	assert.Equal(t, 2, stack.Pop())
	assert.Equal(t, 1, stack.Top())

	stack.Pop()
	assert.Equal(t, nil, stack.Pop())
	assert.Equal(t, 0, stack.Len())

}
