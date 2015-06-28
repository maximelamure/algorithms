package datastructure

import "fmt"

// A Tree is a simple binaray tree
type Tree struct {
	root *TNode
}

type TNode struct {
	key   int
	value int
	left  *TNode
	right *TNode
	level int
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

// DFS traverse the tree in pre-order
func (t *Tree) DFS() <-chan int {
	ch := make(chan int)
	go func() {
		if t.root != nil {
			stack := NewStackArray()
			stack.Push(t.root)
			for {
				if stack.IsEmpty() {
					break
				}
				n := stack.Pop().(*TNode)
				ch <- n.value
				if n.right != nil {
					stack.Push(n.right)
				}
				if n.left != nil {
					stack.Push(n.left)
				}
			}
		}
		close(ch)
	}()
	return ch
}

// InOrder traverse the tree in in-order
func (t *Tree) InOrder() <-chan int {
	ch := make(chan int)
	go func() {
		t.inOrder(t.root, ch)
		close(ch)
	}()
	return ch
}

func (t *Tree) inOrder(node *TNode, ch chan int) {
	if node != nil {
		t.inOrder(node.left, ch)
		ch <- node.value
		t.inOrder(node.right, ch)
	}
}

// PostOrder traverse the tree in post-order
func (t *Tree) PostOrder() <-chan int {
	ch := make(chan int)
	go func() {
		t.postOrder(t.root, ch)
		close(ch)
	}()
	return ch
}

func (t *Tree) postOrder(node *TNode, ch chan int) {
	if node != nil {
		t.postOrder(node.left, ch)
		t.postOrder(node.right, ch)
		ch <- node.value
	}
}

func (t *Tree) DisplayBFS() {
	if t.root == nil {
		return
	}
	queue := NewQueueLinkedList()
	currentLevel := 0
	nbItemByLevel := 0
	queue.Enqueue(t.root)
	for {
		if queue.IsEmpty() {
			return
		}
		current := queue.Dequeue().(*TNode)
		if currentLevel != current.level {
			fmt.Println("")
			currentLevel++
			nbItemByLevel = 0
		}
		if nbItemByLevel > 0 {
			fmt.Print(",")
		}
		fmt.Print(current.value)
		nbItemByLevel++

		if current.left != nil {
			current.left.level = current.level + 1
			queue.Enqueue(current.left)
		}
		if current.right != nil {
			current.right.level = current.level + 1
			queue.Enqueue(current.right)
		}
	}
}
