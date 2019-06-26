package main

//Chapter 2.3, merge sort, implement it by golang
//Author: Raymond Jiang
//Date:   2019-06-25

//Merge sort O(n*lgn)
func MergeSort(arrs []int, p, r int) {
	if p < r-1 {
		q := (p + r) / 2

		MergeSort(arrs, p, q)
		MergeSort(arrs, q, r)
		Merge(arrs, p, q, r)
	}
}
func Merge(arrs []int, p, q, r int) {

	left := make([]int, q-p)
	copy(left, arrs[p:q])

	right := make([]int, r-q)
	copy(right, arrs[q:r])

	j := 0
	i := 0

	for k := p; k < r; k++ {
		if i == len(left) && j < len(right) {
			arrs[k] = right[j]
			j++
			continue
		}
		if j == len(right) && i < len(left) {
			arrs[k] = left[i]
			i++
			continue
		}
		if left[i] < right[j] {
			arrs[k] = left[i]
			i++

		} else {
			arrs[k] = right[j]
			j++
		}

	}

}
