package sort

import "github.com/maximelamure/algorithms/common"

// Shell Sort is a custom insertion Sort
// Move entries more than one position at a time by h-sorting the array
// The worst case number of compares with the 3x+1 increments is O(n^2/3)
type shellSort struct {
}

func NewShellSort() Sorter {
	return &shellSort{}
}

func (s *shellSort) Sort(arr []int) {
	h := 1
	// Compute H with the rule 3x+1
	for {
		if h > len(arr)/3 {
			break
		}
		h = 3*h + 1
	}

	for {
		if h < 1 {
			break
		}

		// insertionSort
		for i := h; i < len(arr); i++ {

			for j := i; j >= h && arr[j] < arr[j-h]; j -= h {
				common.Swap(arr, j, j-h)
			}

		}

		h = h / 3
	}
}
