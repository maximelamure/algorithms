package datastructure

import (
	"strconv"
	"testing"

	"github.com/maximelamure/api/common"
)

func TestQueueLinkedList(t *testing.T) {
	queue := NewQueueLinkedList()
	testQueue(queue, t)
}

func testQueue(queue Queue, t *testing.T) {
	helper := common.Test{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	helper.Assert(t, !queue.IsEmpty(), "The queue should not be empty")
	helper.Assert(t, queue.Size() == 3, "The queue length should be 3")
	queue.Dequeue()
	queue.Dequeue()
	item := queue.Dequeue()
	helper.Assert(t, item.(int) == 3, "The value should be 3. Here: "+strconv.Itoa(item.(int)))
}
