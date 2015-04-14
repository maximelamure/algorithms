package datastructure

import "code.google.com/go/src/pkg/fmt"

// Vertices v and w are connected if there is a path between them.
// Preprocess the Graph to answer queries of the form  is v connected to w in
// constant time.

// A connected component is a maximal set of connected vertices.
type ConnectedComponent struct {
	graph *Graph
	cc    map[int]int
	count int // number of connected component in the graph

}

// NewConnectedComponent preprocess the graph to find connected components in G
func NewConnectedComponent(g *Graph) *ConnectedComponent {
	cc := &ConnectedComponent{
		graph: g,
		cc:    make(map[int]int),
	}
	cc.connect()
	return cc
}

func (c *ConnectedComponent) connect() {
	for v := range c.graph.Adjacencies {
		if _, ok := c.cc[v]; !ok {
			c.dfs(v)
			c.count++
		}
	}
}

func (c *ConnectedComponent) dfs(v int) {
	c.cc[v] = c.count
	for w := range c.graph.Adj(v) {
		if _, ok := c.cc[w]; !ok {
			c.dfs(w)
		}
	}
}

// Connected returns if v an w are connected.
func (c *ConnectedComponent) Connected(v, w int) bool {
	cv, okv := c.cc[v]
	if !okv {
		return false
	}
	cw, okw := c.cc[w]
	if !okw {
		return false
	}
	return cw == cv
}

// Count returns the number of connected components
func (c *ConnectedComponent) Count() int {
	return c.count
}

// getID returns the componenent identifier for v
func (c *ConnectedComponent) GetID(v int) int {
	if v, ok := c.cc[v]; ok {
		return v
	}
	return -1 //FIXME
}

func (c *ConnectedComponent) Display() {
	for key, val := range c.cc {
		fmt.Println(key, " - ", val)
	}
}
