package sort

// mergeSort implements the Merge Sort:
// - Divide array into two values
// - Recusrively sort each half
// - Merge two halves
// Runtime O(nln(n)) average and worst case. Memory: 2N
// Uses at most Nln(N) compares and 6N ln(N) array access to sort
//
// User insertion Sort for small subarrays
// - MergeSort has too much overhead for tiny subarrays
// - Cutoff to insertion for ~ 7 items

type mergeSort struct {
	aux []int
}

func NewMergeSort() Sorter {
	return &mergeSort{}
}

func (s *mergeSort) Sort(arr []int) {
	s.aux = make([]int, len(arr))
	s.mergeSort(arr, 0, len(arr)-1)
}

func (s *mergeSort) mergeSort(arr []int, low, high int) {
	if low < high {
		med := low + (high-low)/2
		s.mergeSort(arr, low, med)
		s.mergeSort(arr, med+1, high)
		if arr[med] < arr[med+1] { // optim. If true, the left and right sub array are already ordered
			return
		}
		s.merge(arr, low, med, high)
	}
}

func (s *mergeSort) merge(arr []int, low, med, heigh int) {
	//copy to the aux array
	for x := low; x <= heigh; x++ {
		s.aux[x] = arr[x]
	}

	// merge logic
	ileft := low
	iright := med + 1

	for i := low; i <= heigh; i++ {
		if ileft > med {
			arr[i] = s.aux[iright]
			iright++
		} else if iright > heigh {
			arr[i] = s.aux[ileft]
			ileft++
		} else if s.aux[ileft] < s.aux[iright] {
			arr[i] = s.aux[ileft]
			ileft++
		} else {
			arr[i] = s.aux[iright]
			iright++
		}
	}
}
