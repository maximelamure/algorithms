package datastructure

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/maximelamure/api/common"
)

func GenerateTree() (tree *Tree, bfs []int, dfs []int, inorder []int, postOrder []int) {
	return &Tree{
			root: &TNode{
				key:   5,
				value: 5,
				left: &TNode{
					key:   2,
					value: 2,
					left: &TNode{
						key:   0,
						value: 0,
					},
					right: &TNode{
						key:   1,
						value: 1,
					},
				},
				right: &TNode{
					key:   8,
					value: 8,
					right: &TNode{
						key:   3,
						value: 3,
					},
				},
			},
		},
		[]int{5, 2, 8, 0, 1, 3},
		[]int{5, 2, 0, 1, 8, 3},
		[]int{0, 2, 1, 5, 8, 3},
		[]int{0, 1, 2, 3, 8, 5}
}

func TestTree(t *testing.T) {
	helper := common.Test{}
	tree, bfs, dfs, inorder, postorder := GenerateTree()
	i := 0
	for x := range tree.BFS() {
		helper.Assert(t, x == bfs[i], "The value should be "+strconv.Itoa(bfs[i])+" here: "+strconv.Itoa(x))
		i++
	}
	i = 0
	fmt.Println("DFS")
	for x := range tree.DFS() {
		log.Print(x, " ")
		helper.Assert(t, x == dfs[i], "The value should be "+strconv.Itoa(dfs[i])+" here: "+strconv.Itoa(x))
		i++
	}

	i = 0
	fmt.Println("InOrder")
	for x := range tree.InOrder() {
		log.Print(x, " ")
		helper.Assert(t, x == inorder[i], "The value should be "+strconv.Itoa(inorder[i])+" here: "+strconv.Itoa(x))
		i++
	}

	i = 0
	fmt.Println("PostOrder")
	for x := range tree.PostOrder() {
		log.Print(x, " ")
		helper.Assert(t, x == postorder[i], "The value should be "+strconv.Itoa(postorder[i])+" here: "+strconv.Itoa(x))
		i++
	}
}
