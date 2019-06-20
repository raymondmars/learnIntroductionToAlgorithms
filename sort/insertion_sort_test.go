package main

import (
	"fmt"
	"gotest.tools/assert"
	"math/rand"
	"strings"
	"testing"
)

func Test_InsertSort(t *testing.T) {
	arrs := []int{1, 3, 4, 2, 8, 5, 6, 7, 9, 0}

	sortedArrs := InsertionSort(arrs)
	assert.Equal(t, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sortedArrs)), ""), "[]"), "0123456789")
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
