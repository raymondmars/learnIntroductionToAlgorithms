package sort

func CountingSort(a []int, b []int, k int) {
	c := make([]int, k)
	for i := 0; i < k; i++ {
		c[i] = 0
	}
	for i := 0; i < len(a); i++ {
		c[a[i]] = c[a[i]] + 1
	}
	for i := 0; i < k; i++ {
		if i > 0 {
			c[i] = c[i] + c[i-1]
		}
	}
	for i := len(a) - 1; i >= 0; i-- {
		b[c[a[i]]-1] = a[i]
		c[a[i]] = c[a[i]] - 1
	}

}
