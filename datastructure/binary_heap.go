package datastructure

import (
	"fmt"

	"github.com/maximelamure/algorithms/common"
)

// A binary heap could be represented by an array
// - Largest key is arr[1], wich is root of binary treee
// - Can use array indices to move through tree
//   - Parent of node k is k/2
//   - Children of node at k are at 2k and 2k+1
// Insert: log N , delMax: log N , max: 1

type BinaryHeap struct {
	arr    []int
	lenght int
}

func NewBinaryHeap(capacity int) *BinaryHeap {
	h := &BinaryHeap{}
	h.arr = make([]int, capacity+1)
	return h
}

// Add add node at the end, them swim it up
// At most 1 + lgN compares
func (b *BinaryHeap) Add(node int) {
	b.lenght++ // start at 1
	b.arr[b.lenght] = node
	b.swim_up(b.lenght)
}

func (b *BinaryHeap) Empty() bool {
	return b.lenght == 0
}

func (b *BinaryHeap) swim_up(k int) {
	for {
		if k > 1 && b.arr[k] > b.arr[k/2] {
			common.Swap(b.arr, k, k/2)
			k = k / 2
		} else {
			break
		}
	}
}

// DeleteMax exchanges root with node at end, them sink it down
// At most 2lnN compares
func (b *BinaryHeap) DeleteMax() {
	common.Swap(b.arr, 1, b.lenght)
	b.lenght--
	b.sink_down(1)
	b.arr[b.lenght+1] = 0
}

func (b *BinaryHeap) sink_down(k int) {
	for {
		if 2*k > b.lenght {
			break
		}
		j := 2 * k

		//Select the higest child
		if j < b.lenght && b.arr[j] < b.arr[j+1] {
			j++
		}
		if b.arr[j] > b.arr[k] {
			common.Swap(b.arr, k, j)
			k = j
		} else {
			break
		}
	}
}

func (b *BinaryHeap) Print() {
	fmt.Println(b.arr)
}
