package apply

import "testing"

func TestTopK(t *testing.T) {
	members := []int{23, 41, 3, 8, 2, 46, 32, 68, 7, 10}
	t.Log(topK(members, 5))
}
