package sort

import (
	"fmt"
	"testing"

	"github.com/maximelamure/algorithms/common"
)

func GetRandomArray() []int {
	return []int{5, 3, 6, 10, 7, 8, 4, 5, 6, 3, 8, 4, 6, 2, 1, 7}
}

func Test_QuickSelect2(t *testing.T) {
	helper := common.Test{}
	array := []int{5, 3, 6, 10, 7, 8, 4, 5, 6}
	helper.Assert(t, QuickSelect(array, 1) == 4, "Error")
}

func Test_Selection(t *testing.T) {
	s := NewSelectionSort()
	testSort(s, t)
}

func Test_Insertion(t *testing.T) {
	s := NewInsertionSort()
	testSort(s, t)
}

func Test_Bubble(t *testing.T) {
	s := NewBubbleSort()
	testSort(s, t)
}

func Test_ShellSort(t *testing.T) {
	s := NewShellSort()
	testSort(s, t)
}

func Test_MergeSort(t *testing.T) {
	s := NewMergeSort()
	testSort(s, t)
}

func Test_MergeButtonSort(t *testing.T) {
	s := NewMergeBottomSort()
	testSort(s, t)
}

func Test_QuickSort(t *testing.T) {
	s := NewQuickSort()
	testSort(s, t)
}

func Test_3WayQuickSort(t *testing.T) {
	s := New3WayQuickSort()
	testSort(s, t)
}

func Test_HeapSort(t *testing.T) {
	helper := common.Test{}
	array := []int{60, 50, 40, 38, 43, 36, 22}
	su := NewShuffle()
	su.Sort(array)
	array = append(array, 0)
	common.Swap(array, 0, len(array)-1)
	fmt.Println(array)
	s := NewHeapSort()
	s.Sort(array)
	fmt.Println(array)
	helper.Assert(t, Sorted(array), "The array is not sorted ")
}

func Test_Shuffle(t *testing.T) {
	helper := common.Test{}
	array := GetRandomArray()
	s := NewInsertionSort()
	s.Sort(array)
	helper.Assert(t, Sorted(array), "The array is not sorted ")
	fmt.Println(array)
	su := NewShuffle()
	su.Sort(array)
	fmt.Println(array)
	helper.Assert(t, !Sorted(array), "The array is sorted ")
}

func testSort(s Sorter, t *testing.T) {
	helper := common.Test{}
	array := GetRandomArray()
	s.Sort(array)
	fmt.Println(array)
	helper.Assert(t, Sorted(array), "The array is not sorted ")
}

func Sorted(arr []int) bool {
	sorted := true
	for i := range arr {
		if i != 0 && arr[i] < arr[i-1] {
			sorted = false
			break
		}
	}
	return sorted
}
