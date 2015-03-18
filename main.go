package main

import "github.com/maximelamure/algorithms/unionfind"

func main() {
	exo1 := unionfind.NewQuickFind(10)
	exo1.Union(1, 5)
	exo1.Union(4, 9)
	exo1.Union(1, 7)
	exo1.Union(0, 2)
	exo1.Union(0, 6)
	exo1.Union(3, 2)
	exo1.String()

	exo2 := unionfind.NewWeightedQU(10)
	exo2.Union(6, 5)
	exo2.Union(2, 0)
	exo2.Union(3, 8)
	exo2.Union(6, 4)
	exo2.Union(4, 9)
	exo2.Union(6, 1)
	exo2.Union(2, 8)
	exo2.Union(1, 3)
	exo2.Union(9, 7)
	exo2.String()
}
