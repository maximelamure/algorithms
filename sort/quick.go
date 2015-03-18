package sort

type quickSort struct{}

func NewQuickSort() Sorter {
	return &quickSort{}
}

func (s *quickSort) Sort(arr []int) {
	su := NewShuffle()
	su.Sort(arr) //Shuffle to avoid the worst case
	s.quickSort(arr, 0, len(arr)-1)
}

func (s *quickSort) quickSort(arr []int, lo, hi int) {
	if hi <= lo {
		return
	}
	p := Partition(arr, lo, hi)
	s.quickSort(arr, lo, p-1)
	s.quickSort(arr, p+1, hi)
}
