package orderstatistics

import "introductionToAlgorithms/sort"

func RandomSelect(list []int, p, r, i int) int {
	if p == r {
		return list[p]
	}
	q := sort.RandPartition(list, p, r)

	k := q - p + 1
	if k == i {
		return list[q]
	}
	if i < k {
		return RandomSelect(list, p, q-1, i)
	} else {
		return RandomSelect(list, q+1, r, i-k)
	}
}
