package sort

import "testing"

func TestRandQuickSort(t *testing.T) {
	list := []int{16, 4, 10, 14, 7, 9, 80, 3, 2, 8, 1}
	RandQuickSort(list, 0, len(list)-1)
	t.Log(list)
}
