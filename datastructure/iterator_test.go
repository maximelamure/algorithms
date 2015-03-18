package datastructure

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/maximelamure/api/common"
)

func TestIterator(t *testing.T) {
	helper := common.Test{}
	i := NewSimpleIterator()
	counter := 0

	for item := range i.Names() {
		fmt.Println(item)
		counter++
	}

	helper.Assert(t, counter == i.Length, "The number of iteration should be "+strconv.Itoa(i.Length)+" . Here: "+strconv.Itoa(counter))
}
