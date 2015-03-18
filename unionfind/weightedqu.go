package unionfind

import "fmt"

// WeightedQU is a QuickUnion to avoid tall trees
// - Keep track of size of each tree
// - Balance by linking root of smaller tree to root of larger tree

type weightedQU struct {
	IDs     []int
	Weights []int
}

// NewWeightedQU initialize a new Weighted QuickUnion structure
func NewWeightedQU(n int) *weightedQU {
	result := &weightedQU{}
	result.IDs = make([]int, n)
	result.Weights = make([]int, n)
	for x := range result.IDs {
		result.IDs[x] = x
		result.Weights[x] = 1
	}
	return result
}

// Root returns the root of r
func (q *weightedQU) Root(r int) int {
	for {
		if r == q.IDs[r] {
			break
		}
		r = q.IDs[r]
	}
	return r
}

// Connectect check if p and r have the same root.
// - true: they are connected
// - false: the are not connected
func (q *weightedQU) Connected(p, r int) bool {
	return q.Root(p) == q.Root(r)
}

// Union merges p and r by linking root of smaller tree to root of larger tree
func (q *weightedQU) Union(p, r int) {
	qr := q.Root(r)
	qp := q.Root(p)
	if qr == qp {
		return
	}

	if q.Weights[qr] > q.Weights[qp] {
		q.IDs[qp] = qr
		q.Weights[qr] += q.Weights[qp]
	} else {
		q.IDs[qr] = qp
		q.Weights[qp] += q.Weights[qr]
	}

}

func (q *weightedQU) String() {
	fmt.Println("WeightedQU")
	for x := range q.IDs {
		fmt.Print(q.IDs[x], " ")
	}
	fmt.Println(" ")
}
