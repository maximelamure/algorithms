package datastructure

// A Tree is a simple binaray tree
type Tree struct {
	root *TNode
}

type TNode struct {
	key   int
	value int
	left  *TNode
	right *TNode
}

// BFS traverses the tree in the level order
func (t *Tree) BFS() <-chan int {
	ch := make(chan int)
	go func() {
		if t.root != nil {
			queue := NewQueueLinkedList()
			queue.Enqueue(t.root)
			for {
				if queue.IsEmpty() {
					break
				}
				n := queue.Dequeue().(*TNode)
				ch <- n.value
				if n.left != nil {
					queue.Enqueue(n.left)
				}
				if n.right != nil {
					queue.Enqueue(n.right)
				}
			}
		}
		close(ch)
	}()

	return ch
}
