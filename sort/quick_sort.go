package sort

func QuickSort(list []int, start, end int) {
	if len(list) == 0 || start < 0 || start > end || end >= len(list) {
		return
	}
	q := Partition(list, start, end)
	QuickSort(list, start, q-1)
	QuickSort(list, q+1, end)
}

func Partition(list []int, p, r int) int {
	i := p - 1
	for j := p; j < r; j++ {
		if list[j] <= list[r] {
			i = i + 1
			temp := list[j]
			list[j] = list[i]
			list[i] = temp
		}
	}
	temp := list[r]
	list[r] = list[i+1]
	list[i+1] = temp
	return i + 1
}
