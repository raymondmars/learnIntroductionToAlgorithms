package main

//Chapter 2.1, insertion sort, implement it by golang
//Author: Raymond Jiang
//Date:   2019-06-20

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
