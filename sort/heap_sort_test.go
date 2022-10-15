package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeapify(t *testing.T) {
	trees := []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}
	MaxHeapify(trees, 1)
	assert.Equal(t, []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}, trees)
	t.Log("ok")
}

func TestBuildMaxHeapify(t *testing.T) {
	trees := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	BuildMaxHeapify(trees)
	t.Log(trees)
}

func TestBuildMinHeapify(t *testing.T) {
	trees := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	BuildMinHeapify(trees)
	t.Log(trees)
}

func TestHeapSort(t *testing.T) {
	trees := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7, 5}
	HeapSort(trees)
	t.Log(trees)
}
