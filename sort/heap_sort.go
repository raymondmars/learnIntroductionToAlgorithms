package sort

func MaxHeapify(trees []int, target int) {
	if len(trees) == 0 {
		return
	}
	l := ((target + 1) * 2) - 1
	r := (target + 1) * 2
	largest := target
	if l < len(trees) && trees[l] > trees[target] {
		largest = l
	}
	if r < len(trees) && trees[r] > trees[largest] {
		largest = r
	}
	if largest != target {
		temp := trees[target]
		trees[target] = trees[largest]
		trees[largest] = temp
		MaxHeapify(trees, largest)
	}
}

func MinHeapify(trees []int, i int) {
	if len(trees) == 0 {
		return
	}
	l := ((i + 1) * 2) - 1
	r := (i + 1) * 2
	minIndex := i
	if l < len(trees) && trees[l] < trees[minIndex] {
		minIndex = l
	}
	if r < len(trees) && trees[r] < trees[minIndex] {
		minIndex = r
	}
	if i != minIndex {
		temp := trees[i]
		trees[i] = trees[minIndex]
		trees[minIndex] = temp
		MinHeapify(trees, minIndex)
	}
}

func BuildMaxHeapify(trees []int) {
	if len(trees) == 0 {
		return
	}
	size := len(trees)
	for i := (size / 2) - 1; i >= 0; i-- {
		MaxHeapify(trees, i)
	}
}

func BuildMinHeapify(trees []int) {
	if len(trees) == 0 {
		return
	}
	size := len(trees)
	for i := (size / 2) - 1; i >= 0; i-- {
		MinHeapify(trees, i)
	}
}

func HeapSort(trees []int) {
	if len(trees) == 0 {
		return
	}
	BuildMaxHeapify(trees)
	for i := len(trees) - 1; i >= 1; i-- {
		temp := trees[0]
		trees[0] = trees[i]
		trees[i] = temp
		MaxHeapify(trees[:i], 0)
	}
}
