package divideconquer

var fibCache = make(map[int]int)

// 普通递归解法
// Ω(Ψ^n), Ψ=(1+sqrt(5))/2
func Fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

// 从下而上的缓存递归解法
// θ(n)
func FibBottomUp(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if _, ok := fibCache[n]; ok {
		return fibCache[n]
	} else {
		rt := FibBottomUp(n-1) + FibBottomUp(n-2)
		fibCache[n] = rt
		return rt
	}
}
