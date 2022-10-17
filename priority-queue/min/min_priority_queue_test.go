package minpriorityqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	trees := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7, 5}
	queue := NewMinPriorityQueue(trees)
	assert.Equal(t, 1, queue.Min())
}

func TestExtractMin(t *testing.T) {
	trees := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7, 5}
	queue := NewMinPriorityQueue(trees)
	t.Log(trees)
	assert.Equal(t, 1, queue.ExtractMin())
	assert.Equal(t, len(trees)-1, len(queue.QueueData()))
	// t.Log(queue.QueueData())
	assert.Equal(t, 2, queue.Min())
}

func TestDecreaseKey(t *testing.T) {
	trees := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7, 5}
	queue := NewMinPriorityQueue(trees)

	queue.DecreaseKey(9, 2)

	t.Log(queue.QueueData())
	assert.Equal(t, []int{1, 2, 3, 4, 2, 9, 10, 14, 8, 5, 16}, queue.QueueData())
}

func TestInsert(t *testing.T) {
	trees := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7, 5}
	queue := NewMinPriorityQueue(trees)
	queue.Insert(6)
	t.Log(queue.QueueData())
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 10, 14, 8, 7, 16, 9}, queue.QueueData())
}

func TestDelete(t *testing.T) {
	trees := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7, 5}
	queue := NewMinPriorityQueue(trees)
	// t.Log(queue.QueueData())
	queue.Delete(2)

	assert.Equal(t, []int{1, 2, 9, 4, 5, 16, 10, 14, 8, 7}, queue.QueueData())
}
