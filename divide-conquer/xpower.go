package divideconquer

// 普通乘积法求 x^n
// O(n)
func normalXPower(x, n int) int {
	total := 1
	for i := 1; i <= n; i++ {
		total *= x
	}
	return total
}

// 分治法求 x^n
// O(lgn)
func divideConquerXPower(x, n int) int {
	if n == 1 {
		return x
	}
	if n%2 == 0 {
		return divideConquerXPower(x, n/2) * divideConquerXPower(x, n/2)
	} else {
		return divideConquerXPower(x, (n-1)/2) * divideConquerXPower(x, (n-1)/2) * x
	}
}
