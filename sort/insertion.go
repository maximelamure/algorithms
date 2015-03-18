package sort

import "github.com/maximelamure/algorithms/common"

// insertion Sort implement the Insertion Sort
// In iteration i, swap a[i] with each larger entry to its left
// Runtime O(n2) average and worst case. Memory: O(1)

type insertionSort struct {
}

func NewInsertionSort() Sorter {
	return &insertionSort{}
}

func (s *insertionSort) Sort(arr []int) {
	for x := range arr {
		for y := x; y > 0; y-- {
			if arr[y] < arr[y-1] {
				common.Swap(arr, y, y-1)
			} else {
				break
			}

		}
	}
}
