package unionfind

// Path Compression is an improved Weighted QuickUnion
// - Just after computing the root of p, set the id of each examined node to
// point to that root.

type pathCompression struct {
	IDs     []int
	Weights []int
}

// NewPathCompression initialize a new Path Compression structure
func NewPathCompression(n int) *pathCompression {
	result := &pathCompression{}
	result.IDs = make([]int, n)
	result.Weights = make([]int, n)
	for x := range result.IDs {
		result.IDs[x] = x
		result.Weights[x] = 1
	}
	return result
}

// Root returns the root of r
func (q *pathCompression) Root(r int) int {
	for {
		if r == q.IDs[r] {
			break
		}
		q.IDs[r] = q.IDs[q.IDs[r]]
		r = q.IDs[r]
	}
	return r
}

// Connectect check if p and r have the same root.
// - true: they are connected
// - false: the are not connected
func (q *pathCompression) Connected(p, r int) bool {
	return q.Root(p) == q.Root(r)
}

// Union merges p and r by linking root of smaller tree to root of larger tree
func (q *pathCompression) Union(p, r int) {
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
