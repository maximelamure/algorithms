package sort

// QuickSelect returns the kth largest item in the given array
func QuickSelect(a []int, k int) int {
	if k > len(a)-1 {
		return -1
	}

	//Shuffle
	sh := NewShuffle()
	sh.Sort(a)

	i := 0
	j := len(a) - 1

	for {
		if i < j {
			n := Partition(a, i, j)
			if n > k {
				j = n - 1
			} else if n < k {
				i = n + 1
			} else {
				return a[n]
			}
		} else {
			return a[k]
		}
	}
}
