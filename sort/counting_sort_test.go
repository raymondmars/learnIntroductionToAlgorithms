package sort

import "testing"

func TestCountingSort(t *testing.T) {
	a := []int{6, 0, 2, 0, 1, 3, 4, 6, 1, 3, 2}
	b := make([]int, len(a))
	CountingSort(a, b, 7)
	t.Log(b)
}
