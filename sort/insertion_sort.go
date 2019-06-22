package main

//Insertion Sort O(n*n)
func InsertionSort(arrs []int) {
	for i := 1; i < len(arrs); i++ {
		key := arrs[i]
		j := i - 1
		for j >= 0 && arrs[j] > key {
			arrs[j+1] = arrs[j]
			j--
		}
		arrs[j+1] = key
	}
}
