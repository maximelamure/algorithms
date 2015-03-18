package sort

import "github.com/maximelamure/algorithms/common"

// bubleSort implements the Buble Sort
// We start at the begining of the array. and swap the first two elements if
// the first is greater than the second
// Runtime O(n2) average and worst case. Memory: O(1)

type bubbleSort struct {
}

func NewBubbleSort() Sorter {
	return &bubbleSort{}
}

func (s *bubbleSort) Sort(arr []int) {
	for _ = range arr {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				common.Swap(arr, i, i+1)
			}
		}
	}
}
