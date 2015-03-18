package sort

import "github.com/maximelamure/algorithms/common"

// Selection Sort implements the Selection Sort
// In iteration i, Find the smallest element using a linear scan and
// move it to the front (Swap a[i] and a[min])
// Runtime O(n^2) average and worst case. Memory: O(1)

type selectionSort struct {
}

func NewSelectionSort() Sorter {
	return &selectionSort{}
}

func (s *selectionSort) Sort(arr []int) {
	for x := range arr {
		min := x
		for y := x; y < len(arr); y++ {
			if arr[y] < arr[min] {
				min = y
			}
		}
		common.Swap(arr, x, min)
	}
}
