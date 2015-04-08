// In progress
package datastructure

import "fmt"

type DFSPath struct {
	Marked map[int]bool
	Adj    map[int]int
	G      *Graph
}

func NewDFSPath(g *Graph) *DFSPath {
	dfspath := &DFSPath{
		Marked: make(map[int]bool),
		Adj:    make(map[int]int),
		G:      g,
	}
	return dfspath
}

func (g *DFSPath) DFS(v int) {
	g.Marked[v] = true
	for w := range g.G.Adj(v) {
		if !g.Marked[w] {
			g.DFS(w)
			g.Adj[w] = v
		}
	}
}

func (g *DFSPath) Print() {
	fmt.Println("Marked")
	for k, _ := range g.Marked {
		fmt.Println(k)
	}
	fmt.Println("ADJ")
	for k, v := range g.Adj {
		fmt.Println(k, "-", v)
	}
}

func (g *DFSPath) HasPathTo(v int) bool {
	return g.Marked[v]
}

// PathTo return a path between 2 vertiecs
// FIXME: s should be the root. We should not pass it in the parameter
func (g *DFSPath) PathTo(s, v int) <-chan int {
	if !g.HasPathTo(v) {
		return nil
	}
	stack := NewStackArray()

	for x := v; x != s; x = g.Adj[x] {
		stack.Push(x)
	}
	stack.Push(s)
	return stack.Iterate()
}
