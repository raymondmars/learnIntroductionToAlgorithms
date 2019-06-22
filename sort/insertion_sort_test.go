package main

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_InsertSort(t *testing.T) {

	a := []int{1, 3, 4, 2, 8, 5, 6, 7, 9, 0}
	b := make([]int, 10)
	copy(b, a)

	sort.Ints(a)
	InsertionSort(b)

	assert.Equal(t, a, b)

}
func Benchmark_InsertSort(b *testing.B) {
	b.StopTimer()
	arrs := []int{}
	for i := 0; i < b.N; i++ {
		arrs = append(arrs, rand.Intn(b.N))
	}

	b.StartTimer()
	InsertionSort(arrs)
}
