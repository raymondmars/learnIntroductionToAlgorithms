package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCutRod(t *testing.T) {
	priceList := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	maxValue := CutRod(priceList, 4)
	// t.Log(maxValue)
	assert.Equal(t, 10, maxValue)
	assert.Equal(t, 18, CutRod(priceList, 7))
	assert.Equal(t, 30, CutRod(priceList, 10))
}

func TestMemoizedCutRod(t *testing.T) {
	priceList := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	maxValue := MemoizedCutRod(priceList, 4)
	// t.Log(maxValue)
	assert.Equal(t, 10, maxValue)
	assert.Equal(t, 18, MemoizedCutRod(priceList, 7))
	assert.Equal(t, 30, MemoizedCutRod(priceList, 10))
}

func TestBottomUpCutRod(t *testing.T) {
	priceList := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	maxValue := BottomUpCutRod(priceList, 4)
	// t.Log(maxValue)
	assert.Equal(t, 10, maxValue)
	assert.Equal(t, 18, BottomUpCutRod(priceList, 7))
	assert.Equal(t, 30, BottomUpCutRod(priceList, 10))
}
