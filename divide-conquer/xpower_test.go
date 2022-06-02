package divideconquer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalXPower(t *testing.T) {
	assert.Equal(t, 8, normalXPower(2, 3))
	assert.Equal(t, 27, normalXPower(3, 3))
	assert.Equal(t, 16, normalXPower(4, 2))
}

func TestDivideConquerXPower(t *testing.T) {
	assert.Equal(t, 8, divideConquerXPower(2, 3))
	assert.Equal(t, 27, divideConquerXPower(3, 3))
	assert.Equal(t, 16, divideConquerXPower(4, 2))
}

func Benchmark_NormalXPower(b *testing.B) {
	for i := 1; i < b.N; i++ {
		normalXPower(2, i)
	}
}

func Benchmark_DivideConquerXPower(b *testing.B) {
	for i := 1; i < b.N; i++ {
		divideConquerXPower(2, i)
	}
}
