package divideconquer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	assert.Equal(t, 0, Fib(0))
	assert.Equal(t, 1, Fib(1))
	assert.Equal(t, 1, Fib(2))
	assert.Equal(t, 2, Fib(3))
	assert.Equal(t, 3, Fib(4))
	assert.Equal(t, 5, Fib(5))
	assert.Equal(t, 8, Fib(6))
}

func TestFibBottomUp(t *testing.T) {
	assert.Equal(t, 0, FibBottomUp(0))
	assert.Equal(t, 1, FibBottomUp(1))
	assert.Equal(t, 1, FibBottomUp(2))
	assert.Equal(t, 2, FibBottomUp(3))
	assert.Equal(t, 3, FibBottomUp(4))
	assert.Equal(t, 5, FibBottomUp(5))
	assert.Equal(t, 8, FibBottomUp(6))
}

func Benchmark_Fib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(30)
	}
}

func Benchmark_FibBottomUp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibBottomUp(30)
	}
}
