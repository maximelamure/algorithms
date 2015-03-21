package datastructure

import (
	"testing"

	"github.com/maximelamure/algorithms/sort"
)

func BinarayHeap() []int {
	return []int{60, 50, 40, 38, 43, 36, 22}
}

func TestBinaryHeap(t *testing.T) {
	bh := BinarayHeap()
	su := sort.NewShuffle()
	su.Sort(bh)
	b := NewBinaryHeap(len(bh))
	for _, val := range bh {
		b.Add(val)
	}
	b.Print()
	for _ = range bh {
		b.DeleteMax()
	}
	b.Print()
}
