package sort

import "github.com/maximelamure/algorithms/common"

//  Heap construction uses <= 2N compares and exchanges
// heapSort uses <= 2 N lg N compares and exchanges
// heapSort is optimal for both time and space, but :
// - Inner loop longer than quickSort's
// - Makes poor use of cache memory
// - Not stable
// worst: 2 N lg N , average: 2 N lg N , best: N lg N
// N log N guarantee , in place
type heapSort struct{}

func NewHeapSort() Sorter {
	return &heapSort{}
}

func (h *heapSort) Sort(arr []int) {
	//create a binay heap
	N := len(arr) - 1
	for x := N / 2; x >= 1; x-- {
		h.sink(arr, x, N)
	}
	// Delete the max

	for {
		if N <= 1 {
			break
		}
		common.Swap(arr, 1, N)
		N--
		h.sink(arr, 1, N)
	}
}

func (h *heapSort) sink(arr []int, k, N int) {
	for {
		if k > N/2 {
			break
		}
		j := 2 * k
		if j < N && arr[j] < arr[j+1] {
			j++
		}
		if arr[k] < arr[j] {
			common.Swap(arr, k, j)
			k = j
		} else {
			break
		}
	}
}
