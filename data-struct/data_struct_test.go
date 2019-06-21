package main

import "testing"
import "gotest.tools/assert"

//Test all queue features
func TestQueue(t *testing.T) {
	queue := &Queue{}
	assert.Equal(t, queue.IsEmpty(), true)
	queue.Push(1)
	assert.Equal(t, queue.IsEmpty(), false)
	queue.Push(2)
	assert.Equal(t, queue.Len(), 2)

	assert.Equal(t, queue.Top(), 1)
	assert.Equal(t, queue.Pop(), 1)
	assert.Equal(t, queue.Top(), 2)

	queue.Pop()
	assert.Equal(t, queue.Pop(), nil)
	assert.Equal(t, queue.Len(), 0)
}

//Test all stack features
func TestStack(t *testing.T) {
	stack := &Stack{}
	assert.Assert(t, stack.IsEmpty())
	stack.Push(1)
	assert.Assert(t, !stack.IsEmpty())
	stack.Push(2)
	assert.Equal(t, stack.Len(), 2)

	assert.Equal(t, stack.Top(), 2)
	assert.Equal(t, stack.Pop(), 2)
	assert.Equal(t, stack.Top(), 1)

	stack.Pop()
	assert.Equal(t, stack.Pop(), nil)
	assert.Equal(t, stack.Len(), 0)

}
