package datastructure

import (
	"bytes"
	"fmt"
)

type Graph struct {
	Adjacencies map[int]*Bag
}

// New Graph creates an empty Graph with v vertices
func NewGraph() *Graph {
	g := &Graph{}
	g.Adjacencies = make(map[int]*Bag)
	return g
}

//AddEdge add an edge v-w
func (g *Graph) AddEdge(v, w int) {
	if _, ok := g.Adjacencies[v]; !ok {
		g.Adjacencies[v] = NewBag()
	}
	g.Adjacencies[v].Add(w)

	if _, ok := g.Adjacencies[w]; !ok {
		g.Adjacencies[w] = NewBag()
	}
	g.Adjacencies[w].Add(v)
}

// Adj returns vertices adjacent to v
func (g *Graph) Adj(v int) <-chan int {
	ch := make(chan int)
	go func() {
		if adj, ok := g.Adjacencies[v]; ok {
			for val := range adj.Iterate() {
				ch <- val
			}
		}
		close(ch)
	}()
	return ch
}

// NbVertices returns the nunber of vertices
func (g *Graph) NbVertices() int {
	return len(g.Adjacencies)
}

// NbEdges returns the number of edges
func (g *Graph) NbEdges() int {
	total := 0
	for _, b := range g.Adjacencies {
		total += b.Size()
	}
	return total / 2
}

// String returns the string representation of the Graph
func (g *Graph) String() string {
	var buffer bytes.Buffer
	for v, b := range g.Adjacencies {
		for w := range b.Iterate() {
			buffer.WriteString(fmt.Sprintf("%s - %s", v, w))
		}
	}
	return buffer.String()
}

// Degree computes the degree
func (g *Graph) Degree(v int) int {
	if v, ok := g.Adjacencies[v]; ok {
		return v.Size()
	}
	return 0 //FIXME manage error
}

// MaxDegree computes the max degree
func (g *Graph) MaxDegree() int {
	max := 0
	for v, _ := range g.Adjacencies {
		currentDegree := g.Degree(v)
		if currentDegree > max {
			max = currentDegree
		}
	}
	return max
}

// NumberOfSelfLoops returns the number of Self loops
func (g *Graph) NumberOfSelfLoops() int {
	loop := 0
	for v, b := range g.Adjacencies {
		for w := range b.Iterate() {
			if w == v {
				loop++
			}
		}
	}
	return loop / 2
}
