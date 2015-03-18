package sort

import "github.com/maximelamure/algorithms/common"

type Sorter interface {
	Sort(a []int)
}

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func Partition(arr []int, lo, hi int) int {
	i := lo
	j := hi
	pivot := lo
	for {

		for {
			if i >= hi || arr[i] > arr[pivot] {
				break
			}
			i++
		}
		for {
			if j <= lo || arr[j] < arr[pivot] {
				break
			}
			j--
		}
		if i >= j {
			break
		}

		common.Swap(arr, i, j)
	}
	common.Swap(arr, pivot, j)
	return j
}
