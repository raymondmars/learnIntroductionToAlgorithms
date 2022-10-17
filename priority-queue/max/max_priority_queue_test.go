package maxpriorityqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	trees := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7, 5}
	queue := NewMaxPriorityQueue(trees)
	assert.Equal(t, 16, queue.Max())
}

func TestExtractMax(t *testing.T) {
	trees := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7, 5}
	queue := NewMaxPriorityQueue(trees)
	t.Log(trees)
	assert.Equal(t, 16, queue.ExtractMax())
	assert.Equal(t, len(trees)-1, len(queue.QueueData()))
	// t.Log(queue.QueueData())
	assert.Equal(t, 14, queue.Max())
}

func TestIncreaseKey(t *testing.T) {
	trees := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7, 5}
	queue := NewMaxPriorityQueue(trees)

	queue.IncreaseKey(10, 18)

	assert.Equal(t, []int{18, 16, 10, 8, 14, 9, 3, 2, 4, 1, 5}, queue.QueueData())
}

func TestInsert(t *testing.T) {
	trees := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7, 5}
	queue := NewMaxPriorityQueue(trees)
	queue.Insert(11)
	// t.Log(queue.QueueData())
	assert.Equal(t, []int{16, 14, 11, 8, 7, 10, 3, 2, 4, 1, 5, 9}, queue.QueueData())
}
