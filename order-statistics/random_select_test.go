package orderstatistics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomSelect(t *testing.T) {
	list := []int{23, 12, 45, 1, 3, 8, 14, 5, 19}
	assert.Equal(t, 12, RandomSelect(list, 0, len(list)-1, 5))
	assert.Equal(t, 23, RandomSelect(list, 0, len(list)-1, 8))
}
