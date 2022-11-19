package dp

import "math"

func CutRod(priceList []int, n int) int {
	if n == 0 {
		return 0
	}
	q := math.MinInt
	for i := 1; i <= n; i++ {
		q = max(q, priceList[i]+CutRod(priceList, n-i))
	}

	return q
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func MemoizedCutRod(priceList []int, n int) int {
	cacheList := make([]int, n+1)
	for i := 0; i < len(cacheList); i++ {
		cacheList[i] = math.MinInt
	}

	return MemoizedCutRodAux(priceList, n, cacheList)
}

func MemoizedCutRodAux(priceList []int, n int, cacheList []int) int {
	if n == 0 {
		return 0
	}
	if cacheList[n] >= 0 {
		return cacheList[n]
	}
	q := math.MinInt
	for i := 1; i <= n; i++ {
		q = max(q, priceList[i]+MemoizedCutRodAux(priceList, n-i, cacheList))
	}
	cacheList[n] = q

	return q
}

func BottomUpCutRod(priceList []int, n int) int {
	cacheList := make([]int, n+1)
	// cacheList[0] = 0
	for j := 1; j <= n; j++ {
		q := math.MinInt
		for i := 1; i <= j; i++ {
			q = max(q, priceList[i]+cacheList[j-i])
		}
		cacheList[j] = q
	}
	return cacheList[n]
}
