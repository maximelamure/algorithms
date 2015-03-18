package sort

func QuickSelect(a []int, k int) int {
	//todo:Shuffle
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
