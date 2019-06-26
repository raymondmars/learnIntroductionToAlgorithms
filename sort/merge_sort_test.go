package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MergeSort(t *testing.T) {

	a := []int{1, 3, 4, 2, 8, 5, 6, 7, 9, 0, 10, 18, 14}
	b := make([]int, len(a))
	copy(b, a)

	sort.Ints(a)

	MergeSort(b, 0, len(b))

	assert.Equal(t, a, b)

	fmt.Printf("%v", b)

}

func Benchmark_MergeSort(b *testing.B) {
	b.StopTimer()
	arrs := []int{}
	for i := 0; i < b.N; i++ {
		arrs = append(arrs, rand.Intn(b.N))
	}

	b.StartTimer()
	MergeSort(arrs, 0, len(arrs))
}
