package sort

import "testing"

func TestQuickSort(t *testing.T) {
	list := []int{16, 4, 10, 14, 7, 9, 80, 3, 2, 8, 1}
	QuickSort(list, 0, len(list)-1)
	t.Log(list)
}
