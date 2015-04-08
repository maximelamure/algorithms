package datastructure

import (
	"strconv"
	"testing"

	"github.com/maximelamure/api/common"
)

func GetGraph() *Graph {
	g := NewGraph()
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 6)
	g.AddEdge(0, 5)
	g.AddEdge(5, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 6)
	g.AddEdge(7, 8)
	g.AddEdge(9, 10)
	g.AddEdge(9, 12)
	g.AddEdge(9, 11)
	g.AddEdge(11, 12)
	g.AddEdge(11, 11)
	return g
}

func TestGraph(t *testing.T) {
	g := GetGraph()
	helper := common.Test{}
	nb := 0
	for _ = range g.Adj(0) {
		nb++
	}
	helper.Assert(t, nb == 4, "The number of adjecent vertices should be 4")

	//nb vertices
	helper.Assert(t, g.NbVertices() == 13, "The number of vertices should be 13")

	//nb edges
	helper.Assert(t, g.NbEdges() == 13, "The number of edges should be 13, get "+strconv.Itoa(g.NbEdges()))

	//degree
	helper.Assert(t, g.Degree(0) == 4, "The degree of the vertice 0 should be 4")

	//max degree
	helper.Assert(t, g.MaxDegree() == 4, "The max degree of the graph should be 4")

	//nb self loop
	helper.Assert(t, g.NumberOfSelfLoops() == 1, "The number of loop  should be 1, get "+strconv.Itoa(g.NumberOfSelfLoops()))
}

func TestDFS(t *testing.T) {
	helper := common.Test{}
	g := GetGraph()
	dfs := NewDFSPath(g)
	dfs.DFS(0)

	nbVertices := 0
	for _ = range dfs.PathTo(0, 5) {
		nbVertices++
	}
	helper.Assert(t, nbVertices == 5, "The number of vertices should be 5, get:"+strconv.Itoa(nbVertices))
	//todo: add more tests
	//nbVertices = 0
	//for _ = range dfs.PathTo(0, 11) {
	//	nbVertices++
	//}
	//helper.Assert(t, nbVertices == 0, "The number of vertices should be 0, get:"+strconv.Itoa(nbVertices))
}
