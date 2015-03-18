package sort

type mergeButtonSort struct {
	aux []int
}

func NewMergeBottomSort() Sorter {
	return &mergeButtonSort{}
}

func (s *mergeButtonSort) Sort(arr []int) {
	s.aux = make([]int, len(arr))
	for sz := 1; sz < len(arr); sz = sz + sz {
		for lo := 0; lo < len(arr)-sz; lo += sz + sz {
			s.merge(arr, lo, lo+sz-1, Min(lo+sz+sz-1, len(arr)-1))
		}
	}
}

func (s *mergeButtonSort) merge(arr []int, low, med, heigh int) {
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
