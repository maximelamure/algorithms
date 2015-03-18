package datastructure

import (
	"testing"

	"github.com/maximelamure/api/common"
)

func TestStackArray(t *testing.T) {
	stack := NewStackArray()
	testStack(stack, t)
}

func TestLinkedList(t *testing.T) {
	stack := NewStackLinkedList()
	testStack(stack, t)
}

func testStack(stack Stack, t *testing.T) {
	helper := common.Test{}
	stack.Push("1")
	stack.Push("2")
	stack.Push("3")
	helper.Assert(t, !stack.IsEmpty(), "The stack should not be empty")
	helper.Assert(t, stack.Size() == 3, "The stack length should be 3")
	stack.Pop()
	stack.Pop()
	pop := stack.Pop()
	helper.Assert(t, pop == "1", "The value should be 1. Here: "+pop)
}
