package unionfind

import "fmt"

// QuickFind is a data structure
// - Interger array id[] of size N
// - Interpretation: p and q are connected if and only if thay have the same id
type quickFind struct {
	IDs []int
}

// NewQuickFind initialize a new QuickFind structure
func NewQuickFind(n int) *quickFind {
	result := &quickFind{}
	result.IDs = make([]int, n)
	for x := range result.IDs {
		result.IDs[x] = x
	}
	return result
}

// Connected check if p and r have the same id.
// - true: they are connected
// - false: they are not connected
func (q *quickFind) Connected(p, r int) bool {
	return q.IDs[p] == q.IDs[r]
}

// Union create an union between p and r. To merge components containing p and q
// change all entries whose id equals id[p] to id[q]
func (q *quickFind) Union(p, r int) {
	pid := q.IDs[p]
	rid := q.IDs[r]
	for x := range q.IDs {
		if q.IDs[x] == pid {
			q.IDs[x] = rid
		}
	}
}

// String prints the IDs array
func (q *quickFind) String() {
	fmt.Println("QuickFind")
	for x := range q.IDs {
		fmt.Print(q.IDs[x], " ")
	}
	fmt.Println(" ")
}
