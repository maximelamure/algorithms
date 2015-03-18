package sort

import "github.com/maximelamure/algorithms/common"

type threeWayQuickSort struct {
}

func New3WayQuickSort() Sorter {
	return &threeWayQuickSort{}
}

func (s *threeWayQuickSort) Sort(arr []int) {
	s.threeWayQuickSort(arr, 0, len(arr)-1)
}

func (s *threeWayQuickSort) threeWayQuickSort(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}

	lt := lo
	gt := hi
	v := arr[lo]
	i := lo

	for {
		if i > gt {
			break
		}
		if v > arr[i] {
			common.Swap(arr, lt, i)
			lt++
			i++
		} else if v < arr[i] {
			common.Swap(arr, i, gt)
			gt--
		} else {
			i++
		}
	}
	s.threeWayQuickSort(arr, lo, lt-1)
	s.threeWayQuickSort(arr, gt+1, hi)

}
