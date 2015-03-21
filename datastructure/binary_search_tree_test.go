package datastructure

import (
	"strconv"
	"testing"

	"github.com/maximelamure/algorithms/common"
)

func TestBST(t *testing.T) {
	helper := common.Test{}
	bst := &binarySearchTree{}

	bst.Put(50, 50)
	bst.Put(25, 25)
	bst.Put(75, 75)
	bst.Put(60, 60)
	bst.Put(90, 90)
	bst.Put(10, 10)
	bst.Put(30, 30)
	bst.String()
	g := bst.Get(30)
	helper.Assert(t, g == 30, "The value should be 30. Here: "+strconv.Itoa(g))
	min := bst.Min()
	helper.Assert(t, min == 10, "The value should be 10. Here: "+strconv.Itoa(min))
	max := bst.Max()
	helper.Assert(t, max == 90, "The value should be 90. Here: "+strconv.Itoa(max))
	floor := bst.Floor(35)
	helper.Assert(t, floor == 30, "The value should be 30. Here: "+strconv.Itoa(floor))
}
