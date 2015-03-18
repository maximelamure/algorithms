package unionfind

// QuickUnion is a data structure
// - Interger array id[] of size N
// - Interpretation: id[i] is parent of  i
// - Root of i is id[id[id[... id[i]...]]]

type quickUnion struct {
	IDs []int
}

// NewQuickUnion initialize a new QuickUnion structure
func NewQuickUnion(n int) *quickUnion {
	result := &quickUnion{}
	result.IDs = make([]int, n)
	for x := range result.IDs {
		result.IDs[x] = x
	}
	return result
}

// Root returns the root of r
func (q *quickUnion) Root(r int) int {
	for {
		if r == q.IDs[r] {
			break
		}
		r = q.IDs[r]
	}
	return r
}

// Root returns the root of r (recursive version)
func (q *quickUnion) RootRec(p int) int {
	for {
		if p == q.IDs[p] {
			return p
		} else {
			return q.RootRec(q.IDs[p])
		}
	}
}

// Connectect check if p and r have the same root.
// - true: they are connected
// - false: the are not connected
func (q *quickUnion) Connected(p, r int) bool {
	return q.Root(p) == q.Root(r)
}

// Union merges p and r  by setting the id of p's root to the id of r's root
func (q *quickUnion) Union(p, r int) {
	q.IDs[q.Root(r)] = q.Root(p)
}
