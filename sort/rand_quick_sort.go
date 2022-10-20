package sort

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandQuickSort(list []int, p, r int) {
	if len(list) > 0 && p < r {
		q := RandPartition(list, p, r)
		RandQuickSort(list, p, q-1)
		RandQuickSort(list, q+1, r)
	}
}

func RandPartition(list []int, p, r int) int {
	i := rand.Intn(r-p+1) + p
	temp := list[r]
	list[r] = list[i]
	list[i] = temp
	return Partition(list, p, r)
}
